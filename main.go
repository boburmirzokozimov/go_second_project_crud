package main

import (
	"github.com/boburmirzokozimov/go_second_project_crud/controllers"
	"github.com/boburmirzokozimov/go_second_project_crud/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.FindPost)
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	router.Run("localhost:8000")
}
