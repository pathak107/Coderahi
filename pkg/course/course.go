package course

import (
	"sort"

	"github.com/gosimple/slug"
	"github.com/pathak107/coderahi-learn/pkg/dto"
	"github.com/pathak107/coderahi-learn/pkg/editorjs"
	"github.com/pathak107/coderahi-learn/pkg/utils"
	"gorm.io/gorm"
)

func CreateCourse(db *gorm.DB, courseDTO *dto.CreateCourseDTO) (uint, error) {
	course := Course{
		Title:     courseDTO.Title,
		Slug:      slug.Make(courseDTO.Title),
		DescShort: courseDTO.DescShort,
		Cost:      courseDTO.Cost,
	}

	result := db.Create(&course)
	if result.Error != nil {
		//TODO: log the result.Error here
		return 0, utils.NewUnexpectedServerError()
	}
	return course.ID, nil
}

//find all courses
func FindAllCourse(db *gorm.DB) ([]Course, error) {
	var courses []Course
	result := db.Find(&courses)
	if result.Error != nil {
		return courses, utils.NewUnexpectedServerError()
	}
	return []Course{}, nil
}

//find course by id
func FindCourseByID(db *gorm.DB, courseID string, queryParams *QuerParamsCourse) (Course, error) {
	var course Course
	var result *gorm.DB
	if queryParams.LoadPosts {
		result = db.Preload("Sections.Posts").Preload("Sections").First(&course, courseID)
	} else if queryParams.LoadSections {
		result = db.Preload("Sections").First(&course, courseID)
	}
	if result.Error != nil {
		return course, utils.NewNotFoundError(ErrCourseNotFound)
	}
	return course, nil
}

func FindCourseBySlug(db *gorm.DB, slug string, queryParams *QuerParamsCourse) (Course, error) {
	var course Course
	var result *gorm.DB
	if queryParams.LoadPosts {
		result = db.Preload("Sections.Posts").Preload("Sections").Where(&Course{Slug: slug}).First(&course)
	} else if queryParams.LoadSections {
		result = db.Preload("Sections").Where(&Course{Slug: slug}).First(&course)
	}
	if result.Error != nil {
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
	course.Cost = courseDTO.Cost
	course.DescMarkdown = utils.ToStringPtr(editorjs.Markdown(courseDTO.DescBody))
	course.DescHTML = utils.ToStringPtr(editorjs.HTML(courseDTO.DescBody))
	course.DescShort = courseDTO.DescShort
	course.DescJson = utils.ToStringPtr(courseDTO.DescBody)

	result := db.Save(&course)
	if result.Error != nil {
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func DeleteCourseByID(db *gorm.DB, courseID string) error {
	result := db.Delete(&Course{}, courseID)
	if result.Error != nil {
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func CreateSection(db *gorm.DB, sectionDTO *dto.CreateSectionDTO) (uint, error) {
	var sections []Section
	result := db.Find(&sections)
	if result.Error != nil {
		//TODO: log the result.Error here
		return 0, utils.NewUnexpectedServerError()
	}
	section := Section{
		Title:       sectionDTO.Title,
		Description: sectionDTO.Description,
		CourseID:    uint(sectionDTO.CourseID),
		Order:       len(sections),
	}

	result = db.Create(&section)
	if result.Error != nil {
		//TODO: log the result.Error here
		return 0, utils.NewUnexpectedServerError()
	}
	return section.ID, nil
}

func EditSectionByID(db *gorm.DB, sectionDTO *dto.EditSectionDTO, sectionID string) error {
	var section Section
	result := db.First(&section, sectionID)
	if result.Error != nil {
		return utils.NewNotFoundError(ErrSectionNotFound)
	}

	section.Title = sectionDTO.Title
	section.Description = sectionDTO.Description

	result = db.Save(&section)
	if result.Error != nil {
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func DeleteSectionByID(db *gorm.DB, sectionID string) error {
	//TODO: Maintain order after  deleteing
	result := db.Delete(&Section{}, sectionID)
	if result.Error != nil {
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func ChangeOrderOfSection(db *gorm.DB, orderDTO *dto.ChangeOrderSectionDTO) error {
	var section Section
	result := db.First(&section, orderDTO.SectionID)
	if result.Error != nil {
		return utils.NewNotFoundError(ErrSectionNotFound)
	}

	var sections []Section
	result = db.Where(&Section{CourseID: section.CourseID}).Find(&sections)
	if result.Error != nil {
		return utils.NewUnexpectedServerError()
	}

	sort.Slice(sections[:], func(i, j int) bool {
		return sections[i].Order < sections[j].Order
	})

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

	result = db.Save(&section)
	if result.Error != nil {
		return utils.NewUnexpectedServerError()
	}
	result = db.Save(&sections)
	if result.Error != nil {
		return utils.NewUnexpectedServerError()
	}

	return nil

}
