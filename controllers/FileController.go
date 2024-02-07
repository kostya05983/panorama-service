package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileController struct{}

func (ctrl FileController) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"fff": ""})
}

func (ctrl FileController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"fff": ""})
}

func (ctrl FileController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ff": "ff"})
}
