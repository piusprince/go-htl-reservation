package api

import (
	"github.com/gofiber/fiber/v2"
	"go-htl-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Pius",
		LastName:  "Prince",
	}
	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user1": "Test User1"})
}
