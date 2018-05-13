package database

import (
	// _ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"link/models"
)

var DB  *gorm.DB
var err error
func init() {
	DB, err = gorm.Open("mysql", "root:123456@/linkstore?charset=utf8")
	if err != nil {
		panic("database connect failed!")
	}

	fmt.Println("database init")


	//	db,err = gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8")
}

func CloseDD(){
	defer DB.Close()
}

// func InsertLinkMap(link LinkMap) (err error) {

// 	return nil
// }

// checkErr(err)
// stmt, err := db.Prepare("INSERT linkmap SET longurl=?, shorturl=?")
// checkErr(err)
// _, err = stmt.Exec("https://blog.csdn.net/wangshubo1989/article/details/7525761421q", "https://75257614")
// checkErr(err)

// defer db.Close()
