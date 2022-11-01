package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pathak107/coderahi-learn/pkg/handler"
	"github.com/pathak107/coderahi-learn/pkg/middleware"
	"github.com/pathak107/coderahi-learn/pkg/services"
)

// @title           Coderahi Learn API
// @version         1.0
// @description     This is the swagger rest api for coderahi learning portal
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	db, err := services.NewDatabaseService()
	if err != nil {
		log.Fatalf("Failed to initialize database")
	}

	h := handler.NewHandler(db)
	// Routes setup
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(static.Serve("/admin", static.LocalFile("./", false)))
	r.Use(static.Serve("/static/public", static.LocalFile("./images", false)))
	r.Use(middleware.ErrorHandler())
	v1 := r.Group("/api/v1")
	{
		post := v1.Group("/post")
		{
			post.GET("", h.FindAllPosts)
			post.GET("/:post_id", h.FindPostByID)
			post.GET("/slug/:slug", h.FindPostBySlug)
			// post.GET("/:slug", h.FindPostBySlug)
			post.POST("", h.CreatePost)
			post.PATCH("/:post_id", h.EditPost)
			post.DELETE("/:post_id", h.DeletePost)
			post.PATCH("/change/order", h.ChangePostOrder)
		}

		course := v1.Group("/course")
		{
			course.GET("", h.FindAllCourses)
			course.GET("/:course_id", h.FindCourseByID)
			course.GET("/slug/:slug", h.FindCourseBySlug)
			course.POST("", h.CreateCourse)
			course.PATCH("/:course_id", h.EditCourse)
			course.DELETE("/:course_id", h.DeleteCourse)

			course.POST("/section", h.CreateSection)
			course.PATCH("/section/:section_id", h.EditSectionByID)
			course.DELETE("/section/:section_id", h.DeleteSectionByID)
			course.PATCH("/section/order", h.ChangeSectionOrder)

			course.POST("/upload/image/:course_id", h.UploadImageHandler)
		}

		cat := v1.Group("/category")
		{
			cat.GET("", h.FindAllCategories)
		}
	}

	r.LoadHTMLFiles("index.html")
	r.GET("/admin/*any", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":8080")
}
