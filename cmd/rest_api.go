package main

import (
	"log"

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
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := services.NewDatabaseService()
	if err != nil {
		log.Fatalf("Failed to initialize database")
	}

	h := handler.NewHandler(db)
	// Routes setup
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Static("/static", "./public")
	r.Use(middleware.ErrorHandler())
	v1 := r.Group("/api/v1")
	{
		v1.Use(middleware.CORSMiddleware())
		post := v1.Group("/post")
		{
			post.Use(middleware.ErrorHandler())
			post.GET("", h.FindAllPosts)
			post.GET("/:post_id", h.FindPostByID)
			// post.GET("/:slug", h.FindPostBySlug)
			post.POST("", h.CreatePost)
			post.PATCH("/:post_id", h.EditPost)
			post.DELETE("/:post_id", h.DeletePost)
		}

		course := v1.Group("/course")
		{
			course.Use(middleware.CORSMiddleware())
			course.GET("", h.FindAllCourses)
			course.GET("/:course_id", h.FindCourseByID)
			// course.GET("/:slug", h.FindCourseBySlug)
			course.POST("", h.CreateCourse)
			course.PATCH("/:course_id", h.EditCourse)
			course.DELETE("/:course_id", h.DeleteCourse)

			course.POST("/section", h.CreateSection)
			course.PATCH("/section/:section_id", h.EditSectionByID)
			course.DELETE("/section/:section_id", h.DeleteSectionByID)
			course.PATCH("/section/order", h.ChangeSectionOrder)
		}
	}

	r.Run(":8080")
}
