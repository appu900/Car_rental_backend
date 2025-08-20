package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/appu900/carrental/model"
	"github.com/appu900/carrental/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("thisisishjasJSagsa")

func RegisterUser(name, phonenumber, password string) (model.User, error) {
	existingUser, err := repository.FindUserByPhoneNumber(phonenumber)
	if err == nil && existingUser.ID != primitive.NilObjectID {
		return model.User{}, fmt.Errorf("user with phone number %s already exists", phonenumber)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	user := model.User{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Phonenumber: phonenumber,
		Password:    string(hashedPassword),
		Role:        "customer",
	}
	_, err = repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func AuthenticateUser(phoneNumber, password string) (string, error) {
	user, err := repository.FindUserByPhoneNumber(phoneNumber)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return "", errors.New("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid password")
	}
	claims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
