package routes

import (
	"github.com/alanaguiar01/go-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRoutes(app *fiber.App, db *gorm.DB) {
	payload := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	app.Post("/users", func(c *fiber.Ctx) error {
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		user := models.User{Name: payload.Name, Email: payload.Email, Password: payload.Password}
		db.Create(&user)
		return c.JSON(user)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var user models.User
		db.Find(&user, "id=?", 2)
		return c.JSON(user)
	})
}
