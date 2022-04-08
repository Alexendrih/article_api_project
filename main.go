package main

import (
	"article_api/database/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome")
}

func setupRoutes(app *fiber.App) {

	app.Get("/api", welcome)
	//user
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	//user_role
	app.Post("/api/useroles", routes.CreateUserole)
	app.Get("/api/useroles", routes.GetUseroles)
	app.Get("/api/useroles/:id", routes.GetUserole)
	app.Put("/api/useroles/:id", routes.UpdateUserole)
	app.Delete("/api/useroles/:id", routes.DeleteUserole)
	//article
	app.Post("/api/articles", routes.CreateArticle)
	app.Get("/api/articles", routes.GetArticles)
	app.Get("/api/articles/:id", routes.GetArticle)
	app.Put("/api/articles/:id", routes.UpdateArticle)
	app.Delete("/api/articles/:id", routes.DeleteArticle)
	//audit_log
	app.Post("/api/audits", routes.CreateAudit)
	app.Get("/api/audits", routes.GetAudits)
	app.Get("/api/audits/:id", routes.GetAudit)
	app.Delete("/api/audits/:id", routes.DeleteAudit)
}
func main() {
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
