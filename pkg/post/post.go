package post

import (
	editorjs "github.com/micheleriva/editorjs-go"
	"github.com/pathak107/coderahi-learn/pkg/utils"
	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB, postDTO *CreatePostDTO) (uint, error) {
	post := Post{
		Title:       postDTO.Title,
		Description: postDTO.Description,
		ImageUrl:    utils.ToStringPtr(postDTO.ImageURL),
		IsBlogPost:  postDTO.IsBlogPost,
		Likes:       0,
		Views:       0,
	}

	result := db.Create(&post)
	if result.Error != nil {
		//TODO: log the result.Error here
		return 0, utils.NewUnexpectedServerError()
	}
	return post.ID, nil
}

func EditPost(db *gorm.DB, postDTO EditPostDTO) error {
	post, err := FindPostByID(db, postDTO.PostID)
	if err != nil {
		return err
	}
	post.Title = postDTO.Title
	post.Description = postDTO.Description
	post.MarkDown = utils.ToStringPtr(editorjs.Markdown(postDTO.Body))
	post.HTMLBody = utils.ToStringPtr(editorjs.HTML(postDTO.Body))
	post.ImageUrl = utils.ToStringPtr(postDTO.ImageURL)
	post.IsBlogPost = postDTO.IsBlogPost

	result := db.Save(&post)
	if result.Error != nil {
		return utils.NewUnexpectedServerError()
	}
	return nil
}

func FindAllPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post
	result := db.Find(&posts)
	if result.Error != nil {
		return posts, utils.NewUnexpectedServerError()
	}
	return posts, nil
}

func FindPostByID(db *gorm.DB, postID uint) (Post, error) {
	var post Post
	result := db.First(&post, postID)
	if result.Error != nil {
		return post, utils.NewNotFoundError(ErrPostNotFound)
	}
	return post, nil
}

func DeletePostByID(db *gorm.DB, postID uint) error {
	result := db.Delete(&Post{}, postID)
	if result.Error != nil {
		return utils.NewUnexpectedServerError()
	}
	return nil
}
