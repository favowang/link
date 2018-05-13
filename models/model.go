package models

import (
	"link/database"
	"github.com/jinzhu/gorm"
)




type LinkMap struct {
	gorm.Model
	LongUrl string `gorm:"primary_key"`
	ShorUrl string
}



func init(){
	database.DB.AutoMigrate(&LinkMap{})
}

func  Add(longurl , shorturl string){
	database.DB.Create(&LinkMap{LongUrl: longurl, ShorUrl: shorturl})
}