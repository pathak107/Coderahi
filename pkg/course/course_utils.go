package course

import (
	"errors"
	"sort"

	"github.com/pathak107/coderahi-learn/pkg/models"
)

var (
	ErrCourseNotFound  = errors.New("course not found")
	ErrSectionNotFound = errors.New("section not found")
)

type QuerParamsCourse struct {
	LoadSections   bool
	LoadPosts      bool
	LoadPostTitles bool
	LoadDrafts     bool
}

func sortSectionsInCourse(course models.Course) models.Course {
	sort.SliceStable(course.Sections, func(i, j int) bool {
		return course.Sections[i].Order < course.Sections[j].Order
	})
	return course
}
