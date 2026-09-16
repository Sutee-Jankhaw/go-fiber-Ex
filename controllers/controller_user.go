package controllers

import (
	"go-fiber-test/database"
	m "go-fiber-test/models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user []m.Users
	result := db.Find(&user)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&user)
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user m.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&user)
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user m.Users
	id := c.Params("id")

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}

func RemoveUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user m.Users

	result := db.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
