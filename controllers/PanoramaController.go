package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PanoramaController struct{}

func (ctrl PanoramaController) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"ff": "f"})
}

func (ctrl PanoramaController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"fff": "ff"})
}

func (ctrl PanoramaController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"fff": "ttt"})
}
