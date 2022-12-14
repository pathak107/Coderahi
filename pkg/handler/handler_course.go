package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pathak107/coderahi-learn/pkg/course"
	"github.com/pathak107/coderahi-learn/pkg/dto"
	"github.com/pathak107/coderahi-learn/pkg/models"
)

func (h *Handler) FindAllCourses(ctx *gin.Context) {
	query := &course.QuerParamsCourse{}
	if ctx.DefaultQuery("section", "false") == "true" {
		query.LoadSections = true
	}
	if ctx.DefaultQuery("post", "false") == "true" {
		query.LoadSections = true
		query.LoadPosts = true
	}

	if ctx.DefaultQuery("draft", "false") == "true" {
		query.LoadDrafts = true
	}

	cacheKey := course.CacheKeyMaker(query, false, false)
	cacheResp := h.cache.Get(cacheKey)
	if cacheResp.Found {
		logger.Printf("Cache hit for %v\n", cacheKey)
		courses := cacheResp.Data.([]models.Course)
		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"courses": courses,
			},
		})
		return
	}

	logger.Printf("Cache miss for %v, querying from db\n", cacheKey)
	courses, err := course.FindAllCourse(h.db, query)
	if err != nil {
		ctx.Error(err)
		return
	}
	h.cache.Put(cacheKey, courses)
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
		if ctx.DefaultQuery("post", "false") == "true" {
			query.LoadPosts = true
		}
	}

	if ctx.DefaultQuery("draft", "false") == "true" {
		query.LoadDrafts = true
	}

	cacheKey := course.CacheKeyMaker(query, false, true) + courseID
	cacheResp := h.cache.Get(cacheKey)
	if cacheResp.Found {
		logger.Printf("Cache hit for %v\n", cacheKey)
		course := cacheResp.Data.(models.Course)
		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"course": course,
			},
		})
		return
	}

	logger.Printf("Cache miss for %v, querying from db\n", cacheKey)
	course, err := course.FindCourseByID(h.db, courseID, query)
	if err != nil {
		ctx.Error(err)
		return
	}
	h.cache.Put(cacheKey, course)

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
		if ctx.DefaultQuery("post", "false") == "true" {
			query.LoadPosts = true
		}
	}

	if ctx.DefaultQuery("draft", "false") == "true" {
		query.LoadDrafts = true
	}

	cacheKey := course.CacheKeyMaker(query, true, false) + slug

	cacheResp := h.cache.Get(cacheKey)
	if cacheResp.Found {
		logger.Printf("Cache hit for %v\n", cacheKey)
		course := cacheResp.Data.(models.Course)
		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"course": course,
			},
		})
		return
	}

	logger.Printf("Cache miss for %v, querying from db\n", cacheKey)
	course, err := course.FindCourseBySlug(h.db, slug, query)
	if err != nil {
		ctx.Error(err)
		return
	}
	h.cache.Put(cacheKey, course)

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

	defer h.cache.Purge()
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

	defer h.cache.Purge()

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

	defer h.cache.Purge()

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

	defer h.cache.Purge()

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

	defer h.cache.Purge()

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

	defer h.cache.Purge()

	ctx.JSON(http.StatusOK, gin.H{
		"data": "section saved successfully",
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

	defer h.cache.Purge()

	ctx.JSON(http.StatusOK, gin.H{
		"data": "changed section order successfully",
	})
}
