package main

import (
	// "fmt"
	"link/controllers"
	// "os"
	// "path/filepath"
	// "strings"
	// "database/sql"
	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	//"net/http"
	. "link/database"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("views/**/*")

	router.GET("/", controllers.GetIndex)
	router.POST("/getShortUrl", controllers.GetShortUrl)
	router.POST("/getLongUrl", controllers.GetLongUrl)
	router.Static("/static", "./static")
	defer DB.Close()
	router.Run(":8080")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// func getCurrentDirectory() string {
// 	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

// 	return strings.Replace(dir, "\\", "/", -1)
// }
