package routes

import (
	"article_api/database"
	"article_api/models"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uint   `json:"user_id" gorm:"primaryKey"`
	UserNAme   string `json:"login_username"`
	Password   string `json:"login_password"`
	UserFnAme  string `json:"user_fname"`
	UserLnAme  string `json:"user_lname"`
	YearLevel  string `json:"year_level"`
	Department string `json:"department"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	RoleId     int    `json:"user_role_id"`
}

func CreateResponseUser(userModel models.User) User {
	hash, _ := bcrypt.GenerateFromPassword([]byte("login_password"), 10)
	return User{ID: userModel.ID, UserNAme: userModel.UserNAme, Password: string(hash),
		UserFnAme: userModel.UserFnAme, UserLnAme: userModel.UserLnAme, YearLevel: userModel.YearLevel,
		Department: userModel.Department, Email: userModel.Email, Phone: userModel.Phone,
		RoleId: userModel.RoleId}

}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("login_password"))
	if err != nil {
		fmt.Println(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	reponseUsers := []User{}
	for _, user := range users {
		reponseUser := CreateResponseUser(user)
		reponseUsers = append(reponseUsers, reponseUser)
	}

	return c.Status(200).JSON(reponseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		UserNAme   string `json:"login_username"`
		Password   string `json:"login_password"`
		UserFnAme  string `json:"user_fname"`
		UserLnAme  string `json:"user_lname"`
		YearLevel  string `json:"year_level"`
		Department string `json:"department"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.UserNAme = updateData.UserNAme
	user.Password = updateData.Password
	user.UserFnAme = updateData.UserFnAme
	user.UserLnAme = updateData.UserLnAme
	user.YearLevel = updateData.YearLevel
	user.Department = updateData.Department
	user.Email = updateData.Email
	user.Phone = updateData.Phone

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted")

}
