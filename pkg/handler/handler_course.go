package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pathak107/coderahi-learn/pkg/course"
	"github.com/pathak107/coderahi-learn/pkg/dto"
)

func (h *Handler) FindAllCourses(ctx *gin.Context) {
	courses, err := course.FindAllCourse(h.db)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"courses": courses,
		},
	})
}

func (h *Handler) FindCourseByID(ctx *gin.Context) {
	courseID := ctx.Param("course_id")
	query := &course.QuerParamsCourse{}
	if ctx.DefaultQuery("section", "false") == "true" {
		query.LoadSections = true
		if ctx.DefaultQuery("subsection", "false") == "true" {
			query.LoadSectionsAndSubsections = true
		}
	}

	course, err := course.FindCourseByID(h.db, courseID, query)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"course": course,
		},
	})
}

func (h *Handler) FindCourseBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	query := &course.QuerParamsCourse{}
	if ctx.DefaultQuery("section", "false") == "true" {
		query.LoadSections = true
		if ctx.DefaultQuery("subsection", "false") == "true" {
			query.LoadSectionsAndSubsections = true
		}
	}

	course, err := course.FindCourseBySlug(h.db, slug, query)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"course": course,
		},
	})
}

func (h *Handler) CreateCourse(ctx *gin.Context) {
	var createCourseDTO dto.CreateCourseDTO
	if err := ctx.ShouldBind(&createCourseDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	courseID, err := course.CreateCourse(h.db, &createCourseDTO)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"course_id": courseID,
		},
	})
}

func (h *Handler) DeleteCourse(ctx *gin.Context) {
	courseID := ctx.Param("course_id")
	err := course.DeleteCourseByID(h.db, courseID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "course deleted succesfully",
	})
}

func (h *Handler) EditCourse(ctx *gin.Context) {
	var editCourseDTO dto.EditCourseDTO
	courseID := ctx.Param("course_id")
	if err := ctx.ShouldBind(&editCourseDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := course.EditCourseByID(h.db, &editCourseDTO, courseID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "course saved successfully",
	})
}

func (h *Handler) CreateSection(ctx *gin.Context) {
	var createSectionDTO dto.CreateSectionDTO
	if err := ctx.ShouldBind(&createSectionDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sectionID, err := course.CreateSection(h.db, &createSectionDTO)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"section_id": sectionID,
		},
	})
}

func (h *Handler) DeleteSectionByID(ctx *gin.Context) {
	sectionID := ctx.Param("section_id")
	err := course.DeleteSectionByID(h.db, sectionID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "section deleted succesfully",
	})
}

func (h *Handler) EditSectionByID(ctx *gin.Context) {
	var editSectionDTO dto.EditSectionDTO
	sectionID := ctx.Param("section_id")
	if err := ctx.ShouldBind(&editSectionDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := course.EditSectionByID(h.db, &editSectionDTO, sectionID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "section saved successfully",
	})
}

func (h *Handler) CreateSubsection(ctx *gin.Context) {
	var createSubsectionDTO dto.CreateSubsectionDTO
	if err := ctx.ShouldBind(&createSubsectionDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	subsectionID, err := course.CreateSubsection(h.db, &createSubsectionDTO)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"subsection_id": subsectionID,
		},
	})
}

func (h *Handler) DeleteSubsectionByID(ctx *gin.Context) {
	subsectionID := ctx.Param("subsection_id")
	err := course.DeleteSubsectionByID(h.db, subsectionID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "subsection deleted succesfully",
	})
}

func (h *Handler) EditSubsectionByID(ctx *gin.Context) {
	var editSubsectionDTO dto.EditSubsectionDTO
	subsectionID := ctx.Param("section_id")
	if err := ctx.ShouldBind(&editSubsectionDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := course.EditSubsectionByID(h.db, &editSubsectionDTO, subsectionID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "subssection saved successfully",
	})
}

func (h *Handler) ChangeSectionOrder(ctx *gin.Context) {
	var changeOrderDTO dto.ChangeOrderSectionDTO
	if err := ctx.ShouldBind(&changeOrderDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := course.ChangeOrderOfSection(h.db, &changeOrderDTO)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": "changed section order successfully",
	})
}

func (h *Handler) ChangeSubsectionOrder(ctx *gin.Context) {
	var changeOrderDTO dto.ChangeOrderSubsectionDTO
	if err := ctx.ShouldBind(&changeOrderDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := course.ChangeOrderOfSubsection(h.db, &changeOrderDTO)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": "changed subsection order successfully",
	})
}
