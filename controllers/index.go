package controllers

import (
	"github.com/gin-gonic/gin"
	"link/models"
	"net/http"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index.html", gin.H{})
}

// func GetShortUrl(c *gin.Context) {
// 	longurl := c.Param("longurl")
// 	shorturl := models.GetShortUrl(longurl)
// 	c.String(http.StatusOK, shorturl)
// }
func PostShortUrl(c *gin.Context) {
	longurl := c.PostForm("longurl")

	shorturl := models.GetShortUrl(longurl)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"longUrl":  longurl,
		"shortUrl": shorturl,
	})
}

func GetLongUrl(c *gin.Context) {
	shorturl := c.Param("shorturl")

	longurl, _ := models.GetLongUrl(shorturl)

	c.String(http.StatusOK, longurl)

}

func PostLongUrl(c *gin.Context) {
	shorturl := c.PostForm("shorturl")

	longurl, _ := models.GetLongUrl(shorturl)

	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"longUrl":  longurl,
		"shortUrl": shorturl,
	})

}
