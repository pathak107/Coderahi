package course

import (
	"errors"

	"github.com/pathak107/coderahi-learn/pkg/post"
	"gorm.io/gorm"
)

var (
	ErrCourseNotFound     = errors.New("course not found")
	ErrSectionNotFound    = errors.New("section not found")
	ErrSubsectionNotFound = errors.New("subsection not found")
)

type Course struct {
	gorm.Model
	Title        string
	DescHTML     *string
	DescMarkdown *string
	Slug         string
	Cost         int
	Sections     []Section
	ExpectedTime int
	Likes        int
	Views        int
	//Comments
}

type CreateCourseDTO struct {
	Title    string `json:"title"`
	DescBody string `json:"desc_body"`
	Cost     int    `json:"cost"`
}

type EditCourseDTO struct {
	CourseID uint   `json:"course_id"`
	Title    string `json:"title"`
	DescBody string `json:"desc_body"`
	Cost     int    `json:"cost"`
}

type Section struct {
	gorm.Model
	CourseID     uint
	Title        string
	Description  string
	ExpectedTime int //in minutes
	Order        int
	Subsections  []Subsection
}

type CreateSectionDTO struct {
	CourseID    int    `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
}

type EditSectionDTO struct {
	SectionID   int    `json:"section_id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
}

type Subsection struct {
	gorm.Model
	SectionID    uint
	Title        string
	Description  string
	ExpectedTime int //in minutes
	PostID       uint
	Order        int
	Post         post.Post
	// Video
}

type CreateSubsectionDTO struct {
	SectionID   uint   `json:"section_id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
}

type EditSubsectionDTO struct {
	SectionID    uint   `json:"section_id"`
	SubsectionID uint   `json:"subsection_id"`
	Title        string `json:"title"`
	Description  string `json:"desc"`
	PostID       uint   `json:"post_id"`
}

type ChangeOrderSectionDTO struct {
	SectionID uint `json:"section_id"`
	Order     int  `json:"order"`
}

type ChangeOrderSubsectionDTO struct {
	SubsectionID uint `json:"subsection_id"`
	Order        int  `json:"order"`
}

type AddPostSubsectionDTO struct {
	post.CreatePostDTO
	SubsectionID uint `json:"subsection_id"`
}

type QuerParamsCourse struct {
	LoadSections               bool
	LoadSectionsAndSubsections bool
}
