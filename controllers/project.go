package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/noclueps/fablab/models"
	"github.com/noclueps/fablab/utils"
)

func CreateProject(c *fiber.Ctx) error {
	tokenStr := c.Cookies("jwt")
	_, err := utils.ExtractClaims(tokenStr)

	if !err {
		return fiber.NewError(fiber.StatusBadRequest, "Unauthorized")
	}

	c.JSON(models.Project{
		Images: []string{"something", "something_else"},
	})
	
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