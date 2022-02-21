package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {
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