package handler

import (
	"context"
	"time"

	"github.com/appu900/carrental/config"
	"github.com/appu900/carrental/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterAdmin(c *fiber.Ctx) error {
	adminCollection := config.GetCollection("admins")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var body struct {
		Name        string `json:"name"`
		Password    string `json:"password"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	}
	var existingAdmin model.Admin
	err := adminCollection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&existingAdmin)
	if err == nil && existingAdmin.ID != primitive.NilObjectID {
		return c.Status(400).JSON(fiber.Map{
			"error": "Admin with this email already exists",
		})
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid Request",
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	newAdmin := model.Admin{
		ID:          primitive.NewObjectID(),
		Name:        body.Name,
		Password:    string(hashedPassword),
		PhoneNumber: body.PhoneNumber,
		Email:       body.Email,
		Role:        "Admin",
	}

	_, err = adminCollection.InsertOne(ctx, newAdmin)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to register admin",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"success":  true,
		"admin_id": newAdmin.ID.Hex(),
	})
}

func LoginAdmin(c *fiber.Ctx) error {
	adminCollection := config.GetCollection("admins")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var body struct {
		Name        string `json:"name"`
		Password    string `json:"password"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	}
	var admin model.Admin
	err := adminCollection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&admin)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	if bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password)) != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	claims := jwt.MapClaims{
		"admin_id": admin.ID.Hex(),
		"role":     admin.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("thisisishjasJSagsa"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success":  true,
		"message":  "Login successfull",
		"token":    signedToken,
		"admin_id": admin.ID.Hex(),
	})

}
