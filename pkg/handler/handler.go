package handler

import (
	"encoding/json"
	"strconv"

	"github.com/pathak107/coderahi-learn/pkg/course"
	"github.com/pathak107/coderahi-learn/pkg/post"
	"github.com/pathak107/coderahi-learn/pkg/utils"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(d *gorm.DB) *Handler {
	return &Handler{
		db: d,
	}
}

func (h *Handler) FindAllPosts() (string, error) {
	posts, err := post.FindAllPosts(h.db)
	if err != nil {
		return "", err
	}
	return NewDataRespHandler("posts", posts)
}

func (h *Handler) FindPostByID(postID string) (string, error) {
	pID, _ := strconv.Atoi(postID)
	post, err := post.FindPostByID(h.db, uint(pID))
	if err != nil {
		return "", err
	}
	return NewDataRespHandler("post", post)
}

func (h *Handler) CreatePost(body string) (string, error) {
	var req post.CreatePostDTO
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	postID, err := post.CreatePost(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewDataRespHandler("postID", postID)
}

func (h *Handler) DeletePost(postID string) (string, error) {
	pID, _ := strconv.Atoi(postID)
	err := post.DeletePostByID(h.db, uint(pID))
	if err != nil {
		return "", err
	}
	return NewSuccessDeletionRespHandler()
}

func (h *Handler) EditPost(body string, postID string) (string, error) {
	var req post.EditPostDTO
	pID, _ := strconv.Atoi(postID)
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	req.PostID = uint(pID)
	err = post.EditPost(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewSuccessEditRespHandler()
}

func (h *Handler) FindAllCourses() (string, error) {
	courses, err := course.FindAllCourse(h.db)
	if err != nil {
		return "", err
	}
	return NewDataRespHandler("courses", courses)
}

func (h *Handler) FindCourseByID(courseID string, queryParams map[string]string) (string, error) {
	cID, _ := strconv.Atoi(courseID)
	query := &course.QuerParamsCourse{}
	if queryParams["sections"] == "true" {
		query.LoadSections = true
		if queryParams["subsectionns"] == "true" {
			query.LoadSectionsAndSubsections = true
		}
	}

	course, err := course.FindCourseByID(h.db, uint(cID), query)
	if err != nil {
		return "", err
	}
	return NewDataRespHandler("course", course)
}

func (h *Handler) CreateCourse(body string) (string, error) {
	var req course.CreateCourseDTO
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	err = course.CreateCourse(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewSuccessCreationRespHandler()
}

func (h *Handler) DeleteCourse(courseID string) (string, error) {
	cID, _ := strconv.Atoi(courseID)
	err := course.DeleteCourseByID(h.db, uint(cID))
	if err != nil {
		return "", err
	}
	return NewSuccessDeletionRespHandler()
}

func (h *Handler) EditCourse(body string, postID string) (string, error) {
	var req course.EditCourseDTO
	cID, _ := strconv.Atoi(postID)
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	req.CourseID = uint(cID)
	err = course.EditCourseByID(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewSuccessEditRespHandler()
}

func (h *Handler) CreateSection(body string) (string, error) {
	var req course.CreateSectionDTO
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	err = course.CreateSection(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewSuccessCreationRespHandler()
}

func (h *Handler) DeleteSectionByID(sectionID string) (string, error) {
	sID, _ := strconv.Atoi(sectionID)
	err := course.DeleteSectionByID(h.db, uint(sID))
	if err != nil {
		return "", err
	}
	return NewSuccessDeletionRespHandler()
}

func (h *Handler) EditSectionByID(body string, sectionID string) (string, error) {
	var req course.EditSectionDTO
	sID, _ := strconv.Atoi(sectionID)
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	req.SectionID = sID
	err = course.EditSectionByID(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewSuccessEditRespHandler()
}

func (h *Handler) CreateSubsection(body string) (string, error) {
	var req course.CreateSubsectionDTO
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	postID, err := course.CreateSubsection(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewDataRespHandler("postID", postID)
}

func (h *Handler) DeleteSubsectionByID(sectionID string) (string, error) {
	ssID, _ := strconv.Atoi(sectionID)
	err := course.DeleteSubsectionByID(h.db, uint(ssID))
	if err != nil {
		return "", err
	}
	return NewSuccessDeletionRespHandler()
}

func (h *Handler) EditSubsectionByID(body string, sectionID string) (string, error) {
	var req course.EditSubsectionDTO
	ssID, _ := strconv.Atoi(sectionID)
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	req.SubsectionID = uint(ssID)
	err = course.EditSubsectionByID(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewSuccessEditRespHandler()
}

func (h *Handler) ChangeSectionOrder(body string) (string, error) {
	var req course.ChangeOrderSectionDTO
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	err = course.ChangeOrderOfSection(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewSuccessEditRespHandler()
}

func (h *Handler) ChangeSubsectionOrder(body string) (string, error) {
	var req course.ChangeOrderSubsectionDTO
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	err = course.ChangeOrderOfSubsection(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewSuccessEditRespHandler()
}
