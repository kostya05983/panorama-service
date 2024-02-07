package main

import (
	"github.com/gin-gonic/gin"
	"panorama/controllers"
)

func main() {
	router := gin.Default()

	file := new(controllers.FileController)

	router.POST("/files", file.Create)
	router.GET("/files/:id", file.Get)
	router.GET("/files", file.GetAll)

	router.Run(":8080")
}
