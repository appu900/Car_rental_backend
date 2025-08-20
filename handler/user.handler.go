package handler

import (
	"github.com/appu900/carrental/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var body struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phoneNumber"`
		Password    string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid Request",
		})
	}

	user, err := service.RegisterUser(body.Name, body.PhoneNumber, body.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"user_id": user.ID.Hex(),
	})
}

func LoginUser(c *fiber.Ctx) error {
	var body struct {
		PhoneNumber string `json:"phoneNumber"`
		Password    string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid Request",
		})
	}

	token, err := service.AuthenticateUser(body.PhoneNumber, body.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid phonenumber or password",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"token":   token,
	})

}
