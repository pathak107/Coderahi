package course

import (
	"strconv"

	"github.com/gosimple/slug"
	"github.com/pathak107/coderahi-learn/pkg/dto"
	"github.com/pathak107/coderahi-learn/pkg/editorjs"
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
func FindAllCourse(db *gorm.DB) ([]models.Course, error) {
	var courses []models.Course
	result := db.Find(&courses)
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
	if queryParams.LoadPosts {
		result = db.Preload("Sections.Posts").Preload("Sections").First(&course, courseID)
		course = sortSectionsInCourse(course)
	} else if queryParams.LoadSections {
		result = db.Preload("Sections").First(&course, courseID)
		course = sortSectionsInCourse(course)
	} else {
		result = db.First(&course, courseID)
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

	if courseDTO.DescBody != "" {
		course.DescMarkdown = utils.ToStringPtr(editorjs.Markdown(courseDTO.DescBody))
		course.DescHTML = utils.ToStringPtr(editorjs.HTML(courseDTO.DescBody))
		course.DescJson = utils.ToStringPtr(courseDTO.DescBody)
	}
	course.DescShort = courseDTO.DescShort

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

	if section.Order > orderDTO.Order {
		for i := orderDTO.Order; i <= section.Order-1; i++ {
			sections[i].Order++
		}
	}

	if section.Order < orderDTO.Order {
		for i := orderDTO.Order; i > section.Order; i-- {
			sections[i].Order--
		}
	}

	section.Order = orderDTO.Order
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
