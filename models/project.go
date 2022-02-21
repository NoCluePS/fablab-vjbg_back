package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	WhoWrokedOn string `json:"whoWorkedOn"`
	Supervisor string `json:"supervisor"`
	Author string `json:"author"`
}