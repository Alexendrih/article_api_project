package routes

import (
	"article_api/database"
	"article_api/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Userole struct {
	ID      uint   `json:"user_ role_id" gorm:"primaryKey"`
	Role    string `json:"role"`
	Remarks string `json:"remarks"`
}

func CreateResponseUserole(useroleModel models.Userole) Userole {
	return Userole{ID: useroleModel.ID, Role: useroleModel.Role, Remarks: useroleModel.Remarks}
}

func CreateUserole(c *fiber.Ctx) error {
	var userole models.Userole

	if err := c.BodyParser(&userole); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&userole)
	responseUserole := CreateResponseUserole(userole)

	return c.Status(200).JSON(responseUserole)
}

func GetUseroles(c *fiber.Ctx) error {
	useroles := []models.Userole{}

	database.Database.Db.Find(&useroles)
	reponseUseroles := []Userole{}
	for _, userole := range useroles {
		reponseUserole := CreateResponseUserole(userole)
		reponseUseroles = append(reponseUseroles, reponseUserole)
	}

	return c.Status(200).JSON(reponseUseroles)
}

func findUserole(id int, userole *models.Userole) error {
	database.Database.Db.Find(&userole, "id = ?", id)
	if userole.ID == 0 {
		return errors.New("userole does not exist")
	}
	return nil
}

func GetUserole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var userole models.Userole

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findUserole(id, &userole); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUserole := CreateResponseUserole(userole)

	return c.Status(200).JSON(responseUserole)

}

func UpdateUserole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var userole models.Userole

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findUserole(id, &userole); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUserole struct {
		Role    string `json:"role"`
		Remarks string `json:"remarks"`
	}

	var updateData UpdateUserole

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	userole.Role = updateData.Role
	userole.Remarks = updateData.Remarks

	database.Database.Db.Save(&userole)

	responseUserole := CreateResponseUserole(userole)
	return c.Status(200).JSON(responseUserole)
}

func DeleteUserole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var userole models.Userole

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findUserole(id, &userole); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&userole).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted")

}
