package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/noclueps/fablab/models"
	"github.com/noclueps/fablab/utils"
)

type createProjectRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	WhoWrokedOn string `json:"whoWorkedOn"`
	Supervisor string `json:"supervisor"`
	Images pq.StringArray `gorm:"type:string[]" json:"images"`
}

func CreateProject(c *fiber.Ctx) error {
	tokenStr := c.Cookies("jwt")
	claims, err := utils.ExtractClaims(tokenStr)
	db := database.Database.DB
	req := new(createProjectRequest)

	if !err {
		return fiber.NewError(fiber.StatusBadRequest, "Unauthorized")
	}

	var user models.User
	db.Where("ID = ?", claims["user_id"] ).First(&user)
	
	if user.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "user not found!")
	}

	var project models.Project{}
	project.Title = req.Title
	project.Description = req.Description
	project.WhoWrokedOn = req.WhoWrokedOn
	project.Supervisor = req.Supervisor
	project.Images = req.Images
	prject.Author = user.Name

	c.JSON(project)

	return nil
}

func GetProjects(c *fiber.Ctx) error {
	return nil
}

func GetProject(c *fiber.Ctx) error {
	id := c.Params("id");
	log.Println(id)

	return nil
}