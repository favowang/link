package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"link/models"
	"fmt"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"pageCount": 1,
	})

	fmt.Println("domodel")

	models.Add("123","456")
}
