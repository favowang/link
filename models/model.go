package models

import (
	"link/database"
)

type LinkMap struct {
	gorm.Model
	LongUrl string
	ShorUrl string
}
