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

func GetUserJson(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users

	gen_z := 0
	gen_y := 0
	gen_x := 0
	baby_bloomer := 0
	gi_generation := 0
	db.Find(&users)
	var dataResults []m.UserRes
	for _, user := range users {
		gen := ""
		if user.Age < 24 {
			gen = "GenZ"
			gen_z += 1
		} else if 24 <= user.Age && user.Age <= 41 {
			gen = "GenY"
			gen_y += 1
		} else if 42 <= user.Age && user.Age <= 56 {
			gen = "GenX"
			gen_x += 1
		} else if 57 <= user.Age && user.Age <= 75 {
			gen = "BabyBloomer"
			baby_bloomer += 1
		} else if user.Age > 75 {
			gen = "G.I generation"
			gi_generation += 1
		}
		d := m.UserRes{
			Name:       user.Name,
			Age:        user.Age,
			EmployeeID: user.EmployeeID,
			Generation: gen,
		}
		dataResults = append(dataResults, d)
	}
	r := m.ResultData{
		Data:         dataResults,
		Count:        len(users),
		GenZ:         gen_z,
		GenY:         gen_y,
		GenX:         gen_x,
		BabyBloomer:  baby_bloomer,
		GIgeneration: gi_generation,
	}
	return c.Status(200).JSON(r)
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
