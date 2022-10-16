package post

import (
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

func CreatePost(db *gorm.DB, postDTO *dto.CreatePostDTO) (uint, error) {
	var postsCount int64
	result := db.Model(models.Post{}).Where(&models.Post{SectionID: uint(postDTO.SectionID)}).Count(&postsCount)
	if result.Error != nil {
		logger.Println(result.Error)
		return 0, utils.NewUnexpectedServerError()
	}

	post := models.Post{
		Title:       postDTO.Title,
		Description: postDTO.Description,
		Slug:        slug.Make(postDTO.Title),
		SectionID:   postDTO.SectionID,
		Order:       int(postsCount),
	}

	result = db.Create(&post)
	if result.Error != nil {
		logger.Println(result.Error)
		return 0, utils.NewUnexpectedServerError()
	}
	return post.ID, nil
}

func EditPost(db *gorm.DB, postDTO *dto.EditPostDTO, postID string) error {
	post, err := FindPostByID(db, postID)
	if err != nil {
		return err
	}

	if postDTO.Title != "" {
		post.Title = postDTO.Title
		post.Slug = slug.Make(postDTO.Title)
	}

	if postDTO.Description != "" {
		post.Description = postDTO.Description
	}

	if postDTO.Body != "" {
		post.MarkDown = utils.ToStringPtr(editorjs.Markdown(postDTO.Body))
		post.HTMLBody = utils.ToStringPtr(editorjs.HTML(postDTO.Body))
		post.BodyJson = utils.ToStringPtr(postDTO.Body)
	}

	result := db.Save(&post)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func FindAllPosts(db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post
	result := db.Find(&posts)
	if result.Error != nil {
		logger.Println(result.Error)
		return posts, utils.NewUnexpectedServerError()
	}
	return posts, nil
}

func FindPostByID(db *gorm.DB, postID string) (models.Post, error) {
	var post models.Post
	result := db.First(&post, postID)
	if result.Error != nil {
		logger.Println(result.Error)
		return post, utils.NewNotFoundError(ErrPostNotFound)
	}
	return post, nil
}

func FindPostBySlug(db *gorm.DB, slug string) (models.Post, error) {
	var post models.Post
	result := db.Where(&models.Post{Slug: slug}).First(&post)
	if result.Error != nil {
		logger.Println(result.Error)
		return post, utils.NewNotFoundError(ErrPostNotFound)
	}
	return post, nil
}

func DeletePostByID(db *gorm.DB, postID string) error {
	// maintaining order after deleting
	var post models.Post
	result := db.First(&post, postID)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewNotFoundError(ErrPostNotFound)
	}

	var posts []models.Post
	result = db.Where(&models.Post{SectionID: post.SectionID}).Order("posts.order asc").Find(&posts)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}

	for i := post.Order + 1; i < len(posts); i++ {
		posts[i].Order--
	}
	result = db.Save(&posts)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}

	result = db.Delete(&models.Post{}, postID)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func ChangeOrderOfPost(db *gorm.DB, orderDTO *dto.ChangeOrderPostDTO) error {
	var post models.Post
	result := db.First(&post, orderDTO.PostID)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewNotFoundError(ErrPostNotFound)
	}

	var posts []models.Post
	result = db.Where(&models.Post{SectionID: post.SectionID}).Order("posts.order asc").Find(&posts)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}

	if post.Order > utils.ToInt(orderDTO.Order) {
		for i := utils.ToInt(orderDTO.Order); i <= post.Order-1; i++ {
			posts[i].Order++
		}
	}

	if post.Order < utils.ToInt(orderDTO.Order) {
		for i := utils.ToInt(orderDTO.Order); i > post.Order; i-- {
			posts[i].Order--
		}
	}

	post.Order = utils.ToInt(orderDTO.Order)
	result = db.Save(&posts)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}

	result = db.Save(&post)
	if result.Error != nil {
		logger.Println(result.Error)
		return utils.NewUnexpectedServerError()
	}

	return nil

}
