package course

import (
	"errors"

	"github.com/pathak107/coderahi-learn/pkg/post"
	"gorm.io/gorm"
)

var (
	ErrCourseNotFound  = errors.New("course not found")
	ErrSectionNotFound = errors.New("section not found")
)

type Course struct {
	gorm.Model
	Title        string
	DescHTML     *string
	DescJson     *string
	DescMarkdown *string
	DescShort    string
	Slug         string
	Cost         int
	Sections     []Section
	ExpectedTime int
	Likes        int
	Views        int
	CategoryID   uint
	//Comments
}

type Section struct {
	gorm.Model
	CourseID     uint
	Title        string
	Description  string
	ExpectedTime int //in minutes
	Order        int
	Posts        []post.Post
}

type QuerParamsCourse struct {
	LoadSections   bool
	LoadPosts      bool
	LoadPostTitles bool
}

type Category struct {
	Name        string
	Slug        string
	Description string
	Courses     []Course
}
