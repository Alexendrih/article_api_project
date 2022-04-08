package routes

import (
	"article_api/database"
	"article_api/models"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Article struct {
	ID         uint   `json:"article_id" gorm:"primaryKey"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Footer     string `json:"footer"`
	CreatedAt  time.Time
	Remarks    string `json:"article_remarks"`
	Status     string `json:"date_approved"`
	ApprovedBy string `json:"approved_by"`
	UserId     int    `json:"user_id"`
}

func CreateResponseArticle(articleModel models.Article) Article {
	return Article{ID: articleModel.ID, Title: articleModel.Title, Body: articleModel.Body,
		Footer: articleModel.Footer, CreatedAt: articleModel.CreatedAt, Remarks: articleModel.Remarks,
		Status: articleModel.Status, ApprovedBy: articleModel.ApprovedBy, UserId: articleModel.UserId}
}

func CreateArticle(c *fiber.Ctx) error {
	var article models.Article

	if err := c.BodyParser(&article); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&article)
	responseArticle := CreateResponseArticle(article)

	return c.Status(200).JSON(responseArticle)
}

func GetArticles(c *fiber.Ctx) error {
	articles := []models.Article{}

	database.Database.Db.Find(&articles)
	reponseArticles := []Article{}
	for _, article := range articles {
		reponseArticle := CreateResponseArticle(article)
		reponseArticles = append(reponseArticles, reponseArticle)
	}

	return c.Status(200).JSON(reponseArticles)
}

func findArticle(id int, article *models.Article) error {
	database.Database.Db.Find(&article, "id = ?", id)
	if article.ID == 0 {
		return errors.New("article does not exist")
	}
	return nil
}

func GetArticle(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var article models.Article

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findArticle(id, &article); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseArticle := CreateResponseArticle(article)

	return c.Status(200).JSON(responseArticle)

}

func UpdateArticle(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var article models.Article

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findArticle(id, &article); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateArticle struct {
		Title   string `json:"title"`
		Body    string `json:"body"`
		Footer  string `json:"footer"`
		Remarks string `json:"article_remarks"`
		Status  string `json:"date_approved"`
	}

	var updateData UpdateArticle

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	article.Title = updateData.Title
	article.Body = updateData.Body
	article.Footer = updateData.Footer
	article.Remarks = updateData.Remarks
	article.Status = updateData.Status

	database.Database.Db.Save(&article)

	responseArticle := CreateResponseArticle(article)
	return c.Status(200).JSON(responseArticle)
}

func DeleteArticle(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var article models.Article

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findArticle(id, &article); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&article).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted")

}
