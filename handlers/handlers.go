package handlers

import (
	"basic-jwt-auth/config"
	"basic-jwt-auth/models"
	"basic-jwt-auth/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"time"
)

func Login(c *fiber.Ctx) error {

	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	day := time.Hour * 24

	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"fav":   user.FavoritePhrase,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(models.LoginResponse{
		Token: t,
	})
}

func Protected(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	favPhrase := claims["fav"].(string)
	return c.SendString("Welcome 👋" + email + " " + favPhrase)
}
