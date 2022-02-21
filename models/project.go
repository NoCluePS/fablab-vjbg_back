package models

import (
	pq "github.com/lib/pq"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	WhoWrokedOn string `json:"whoWorkedOn"`
	Supervisor string `json:"supervisor"`
	Images pq.StringArray `gorm:"type:string[]" json:"images"`
	Author string `json:"author"`
}