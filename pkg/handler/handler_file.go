package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pathak107/coderahi-learn/pkg/course"
	"github.com/pathak107/coderahi-learn/pkg/utils"
)

var (
	logger        = utils.NewLogger()
	fileUploadDst = utils.GetImagesDir()
)

func (h *Handler) UploadImageHandler(ctx *gin.Context) {
	courseID := ctx.Param("course_id")
	// Single file
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.Error(utils.NewFileNotUploadedError())
		return
	}

	exts := strings.Split(file.Filename, ".")
	file.Filename = "img-" + strconv.Itoa(int(time.Now().Unix())) + "." + exts[len(exts)-1]
	logger.Println(file.Filename)
	logger.Println(fileUploadDst)
	// Upload the file to specific dst.
	err = ctx.SaveUploadedFile(file, fileUploadDst+file.Filename)
	if err != nil {
		logger.Printf("error while uploading file: %v \n", err)
		ctx.Error(utils.NewUnexpectedServerError())
		return
	}

	course.UpdateCourseImage(h.db, file.Filename, courseID)

	ctx.JSON(http.StatusOK, gin.H{
		"data": "uploaded successfully",
	})
}
