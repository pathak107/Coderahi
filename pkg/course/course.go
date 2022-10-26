package course

import (
	"strconv"

	"github.com/gosimple/slug"
	"github.com/pathak107/coderahi-learn/pkg/dto"
	"github.com/pathak107/coderahi-learn/pkg/models"
	"github.com/pathak107/coderahi-learn/pkg/utils"
	"gorm.io/gorm"
)

var (
	logger = utils.NewLogger()
)

func CreateCourse(db *gorm.DB, courseDTO *dto.CreateCourseDTO) (uint, error) {
	course := models.Course{
		Title:     courseDTO.Title,
		Slug:      slug.Make(courseDTO.Title),
		DescShort: courseDTO.DescShort,
		Cost:      courseDTO.Cost,
	}

	result := db.Create(&course)
	if result.Error != nil {
		logger.Println(result.Error)
		return 0, utils.NewUnexpectedServerError()
	}
	return course.ID, nil
}

//find all courses
func FindAllCourse(db *gorm.DB, queryParams *QuerParamsCourse) ([]models.Course, error) {
	var courses []models.Course
	var result *gorm.DB

	var preloadConditions string
	if !queryParams.LoadDrafts {
		preloadConditions = "Published=true"
	}

	if queryParams.LoadPosts {
		result = db.Preload("Sections.Posts", preloadConditions).
			Preload("Sections", preloadConditions).
			Preload("Categories").
			Where(preloadConditions).
			Find(&courses)
	} else if queryParams.LoadSections {
		result = db.Preload("Sections", preloadConditions).
			Preload("Categories").
			Where(preloadConditions).
			Find(&courses)
	} else {
		result = db.Preload("Categories").
			Where(preloadConditions).
			Find(&courses)
	}

	if result.Error != nil {
		logger.Println(result.Error)
		return courses, utils.NewUnexpectedServerError()
	}
	return courses, nil
}

//find course by id
func FindCourseByID(db *gorm.DB, courseID string, queryParams *QuerParamsCourse) (models.Course, error) {
	var course models.Course
	var result *gorm.DB
	var preloadConditions string
	if !queryParams.LoadDrafts {
		preloadConditions = "Published=true"
	}
	if queryParams.LoadPosts {
		result = db.Preload("Sections.Posts", preloadConditions).
			Preload("Sections", preloadConditions).
			Preload("Categories").
			First(&course, courseID)
		course = sortSectionsInCourse(course)
	} else if queryParams.LoadSections {
		result = db.Preload("Sections", preloadConditions).
			Preload("Categories").
			First(&course, courseID)
		course = sortSectionsInCourse(course)
	} else {
		result = db.Preload("Categories").First(&course, courseID)
	}

	// sorting the sections
	if result.Error != nil {
		logger.Println(result.Error)
		return course, utils.NewNotFoundError(ErrCourseNotFound)
	}
	return course, nil
}

func FindCourseBySlug(db *gorm.DB, slug string, queryParams *QuerParamsCourse) (models.Course, error) {
	var course models.Course
	var result *gorm.DB
	var preloadConditions string
	if !queryParams.LoadDrafts {
		preloadConditions = "Published=true"
	}
	if queryParams.LoadPosts {
		result = db.Preload("Sections.Posts", preloadConditions).
			Preload("Sections", preloadConditions).
			Preload("Categories").
			Where(&models.Course{Slug: slug}).
			First(&course)
		course = sortSectionsInCourse(course)
	} else if queryParams.LoadSections {
		result = db.Preload("Sections", preloadConditions).
			Preload("Categories").
			Where(&models.Course{Slug: slug}).
			First(&course)
		course = sortSectionsInCourse(course)
	} else {
		result = db.Preload("Categories").Where(&models.Course{Slug: slug}).First(&course)
	}

	// sorting the sections
	if result.Error != nil {
		logger.Println(result.Error)
		return course, utils.NewNotFoundError(ErrCourseNotFound)
	}
	return course, nil
}

func EditCourseByID(db *gorm.DB, courseDTO *dto.EditCourseDTO, courseID string) error {
	course, err := FindCourseByID(db, courseID, &QuerParamsCourse{
		LoadSections: false,
		LoadPosts:    false,
	})
	if err != nil {
		return err
	}
	course.Title = courseDTO.Title
	course.Slug = slug.Make(courseDTO.Title)
	course.Cost = courseDTO.Cost
	course.Published = courseDTO.Publish

	if courseDTO.HTML != "" && courseDTO.Markdown != "" {
		course.MarkDown = utils.ToStringPtr(courseDTO.Markdown)
		course.HTML = utils.ToStringPtr(courseDTO.HTML)
	}
	course.DescShort = courseDTO.DescShort
	db.Model(&course).Association("Categories").Clear()
	var cats []models.Category
	for _, newCat := range courseDTO.Categories {
		cats = append(cats, models.Category{
			Name: newCat.Label,
			Slug: slug.Make(newCat.Label),
		})
	}

	course.Categories = cats

	result := db.Save(&course)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func DeleteCourseByID(db *gorm.DB, courseID string) error {
	result := db.Delete(&models.Course{}, courseID)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func FindSectionByID(db *gorm.DB, sectionID string) (models.Section, error) {
	var section models.Section
	result := db.First(&section, sectionID)

	if result.Error != nil {
		logger.Println(result.Error)
		return section, utils.NewNotFoundError(ErrSectionNotFound)
	}
	return section, nil
}

func CreateSection(db *gorm.DB, sectionDTO *dto.CreateSectionDTO) (uint, error) {
	// Check if course id is valid
	_, err := FindCourseByID(db, strconv.Itoa(sectionDTO.CourseID), &QuerParamsCourse{LoadSections: false, LoadPosts: false})
	if err != nil {
		return 0, utils.NewNotFoundError(ErrCourseNotFound)
	}

	var sectionCount int64
	result := db.Model(&models.Section{}).Where(&models.Section{CourseID: uint(sectionDTO.CourseID)}).Count(&sectionCount)
	if result.Error != nil {
		logger.Println(result.Error)
		return 0, utils.NewUnexpectedServerError()
	}
	section := models.Section{
		Title:       sectionDTO.Title,
		Description: sectionDTO.Description,
		CourseID:    uint(sectionDTO.CourseID),
		Order:       int(sectionCount),
	}

	result = db.Create(&section)
	if result.Error != nil {
		logger.Println(result.Error)
		return 0, utils.NewUnexpectedServerError()
	}
	return section.ID, nil
}

func EditSectionByID(db *gorm.DB, sectionDTO *dto.EditSectionDTO, sectionID string) error {
	var section models.Section
	result := db.First(&section, sectionID)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewNotFoundError(ErrSectionNotFound)
	}

	section.Title = sectionDTO.Title
	section.Description = sectionDTO.Description
	section.ExpectedTime = sectionDTO.ExpectedTime
	section.Published = sectionDTO.Publish

	result = db.Save(&section)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func DeleteSectionByID(db *gorm.DB, sectionID string) error {
	// maintaining order after deleting
	var section models.Section
	result := db.First(&section, sectionID)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewNotFoundError(ErrSectionNotFound)
	}

	var sections []models.Section
	result = db.Where(&models.Section{CourseID: section.CourseID}).Order("sections.order asc").Find(&sections)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}

	for i := section.Order + 1; i < len(sections); i++ {
		sections[i].Order--
	}
	result = db.Save(&sections)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}

	result = db.Delete(&models.Section{}, sectionID)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func ChangeOrderOfSection(db *gorm.DB, orderDTO *dto.ChangeOrderSectionDTO) error {
	var section models.Section
	result := db.First(&section, orderDTO.SectionID)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewNotFoundError(ErrSectionNotFound)
	}

	var sections []models.Section
	result = db.Where(&models.Section{CourseID: section.CourseID}).Order("sections.order asc").Find(&sections)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}

	if section.Order > utils.ToInt(orderDTO.Order) {
		for i := utils.ToInt(orderDTO.Order); i <= section.Order-1; i++ {
			sections[i].Order++
		}
	}

	if section.Order < utils.ToInt(orderDTO.Order) {
		for i := utils.ToInt(orderDTO.Order); i > section.Order; i-- {
			sections[i].Order--
		}
	}

	section.Order = utils.ToInt(orderDTO.Order)
	result = db.Save(&sections)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}

	result = db.Save(&section)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}

	return nil

}

func UpdateCourseImage(db *gorm.DB, fileName string, courseID string) error {
	course, err := FindCourseByID(db, courseID, &QuerParamsCourse{
		LoadSections: false,
		LoadPosts:    false,
	})
	if err != nil {
		return err
	}

	course.ImageURL = utils.ToStringPtr("static/public/images/" + fileName)

	result := db.Save(&course)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}
	return nil
}
