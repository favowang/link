package models

import (
	//"github.com/jinzhu/gorm"
	"fmt"
	. "link/database"
	"link/utils"
)

type LinkMap struct {
	LongUrl string `gorm:"primary_key"`
	ShorUrl string
}

func init() {
	DB.AutoMigrate(&LinkMap{})
}

func Add(longurl, shorturl string) {
	DB.Create(&LinkMap{LongUrl: longurl, ShorUrl: shorturl})
}

func GetShortUrl(longurl string) string {
	var lin LinkMap
	err := DB.First(&lin, "long_url=?", longurl).Error
	if err == nil {
		return lin.ShorUrl
	}
	shorturls, err := utils.Transform(longurl)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	lin.LongUrl = longurl
	lin.ShorUrl = shorturls[0]
	DB.Create(&lin)
	return lin.ShorUrl

}

func GetLongUrl(shorturl string) (string, error) {
	var lin LinkMap
	err := DB.First(&lin, "shor_url=?", shorturl).Error
	if err == nil {
		return lin.LongUrl, nil
	}
	return "warnning:no long url mapped!", err
}
