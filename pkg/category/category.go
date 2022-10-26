package category

import (
	"errors"

	"github.com/gosimple/slug"
	"github.com/pathak107/coderahi-learn/pkg/dto"
	"github.com/pathak107/coderahi-learn/pkg/models"
	"github.com/pathak107/coderahi-learn/pkg/utils"
	"gorm.io/gorm"
)

var (
	logger         = utils.NewLogger()
	ErrCatNotFound = errors.New("category not found")
)

func CreateCategory(db *gorm.DB, catDTO *dto.CategoryDTO) (uint, error) {
	cat := models.Category{
		Name: catDTO.Name,
		Slug: slug.Make(catDTO.Name),
	}

	result := db.Create(&cat)
	if result.Error != nil {
		logger.Println(result.Error)
		return 0, utils.NewUnexpectedServerError()
	}
	return cat.ID, nil
}

func CreateBulkCategory(db *gorm.DB, catDTO []*dto.CategoryDTO) ([]uint, error) {
	var cats []models.Category
	var catIDs []uint

	for _, dto := range catDTO {
		cats = append(cats, models.Category{
			Name: dto.Name,
			Slug: slug.Make(dto.Name),
		})
	}
	result := db.Create(&cats)
	if result.Error != nil {
		logger.Println(result.Error)
		return catIDs, utils.NewUnexpectedServerError()
	}

	for _, cat := range cats {
		catIDs = append(catIDs, cat.ID)
	}
	return catIDs, nil
}

func EditCategory(db *gorm.DB, catDTO *dto.CategoryDTO, catID string) error {
	cat, err := FindCategoryByID(db, catID)
	if err != nil {
		return err
	}

	cat.Name = catDTO.Name
	cat.Slug = slug.Make(catDTO.Name)

	result := db.Save(&cat)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func FindAllCategories(db *gorm.DB) ([]models.Category, error) {
	var cats []models.Category
	result := db.Find(&cats)
	if result.Error != nil {
		logger.Println(result.Error)
		return cats, utils.NewUnexpectedServerError()
	}
	return cats, nil
}

func FindCategoryByID(db *gorm.DB, catID string) (models.Category, error) {
	var cat models.Category
	result := db.Preload("Courses").First(&cat, catID)
	if result.Error != nil {
		logger.Println(result.Error)
		return cat, utils.NewNotFoundError(ErrCatNotFound)
	}
	return cat, nil
}

func FindCategoryBySlug(db *gorm.DB, slug string) (models.Category, error) {
	var cat models.Category
	result := db.Where(&models.Category{Slug: slug}).First(&cat)
	if result.Error != nil {
		logger.Println(result.Error)
		return cat, utils.NewNotFoundError(ErrCatNotFound)
	}
	return cat, nil
}

func DeleteCategoryByID(db *gorm.DB, catID string) error {
	var cat models.Category
	result := db.First(&cat, catID)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewNotFoundError(ErrCatNotFound)
	}

	result = db.Delete(&models.Category{}, catID)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func DeleteCategoryBulk(db *gorm.DB, catIDs []uint) error {
	result := db.Delete(&models.Category{}, catIDs)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}
	return nil
}
