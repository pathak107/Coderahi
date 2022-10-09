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
	r.Static("/static", "./public")
	r.Use(middleware.ErrorHandler())
	v1 := r.Group("/api/v1")
	{
		post := v1.Group("/post")
		{
			post.GET("/", h.FindAllPosts)
			post.GET("/:post_id", h.FindPostByID)
			post.POST("/", h.CreatePost)
			post.PATCH("/", h.EditPost)
			post.DELETE("/:post_id", h.DeletePost)
		}

		course := v1.Group("/course")
		{
			course.GET("/vm", func(ctx *gin.Context) {})
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/login", func(ctx *gin.Context) {})
			auth.POST("/register", func(ctx *gin.Context) {})
		}
	}

	r.Run(":3000")
}
