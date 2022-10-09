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

type Section struct {
	gorm.Model
	CourseID     uint
	Title        string
	Description  string
	ExpectedTime int //in minutes
	Order        int
	Subsections  []Subsection
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

type QuerParamsCourse struct {
	LoadSections               bool
	LoadSectionsAndSubsections bool
}
