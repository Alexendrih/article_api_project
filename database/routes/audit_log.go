package routes

import (
	"article_api/database"
	"article_api/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Audit struct {
	ID       uint   `json:"audit_log_id" gorm:"primaryKey"`
	Register string `json:"date_register"`
	Indate   string `json:"login_date"`
	Outdate  string `json:"logut_date"`
	Userstat string `json:"user_status"`
	UserId   int    `json:"user_id"`
}

func CreateResponseAudit(auditModel models.Audit) Audit {
	return Audit{ID: auditModel.ID, Register: auditModel.Register, Indate: auditModel.Indate,
		Outdate: auditModel.Outdate, Userstat: auditModel.Userstat, UserId: auditModel.UserId}
}

func CreateAudit(c *fiber.Ctx) error {
	var audit models.Audit

	if err := c.BodyParser(&audit); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&audit)
	responseAudit := CreateResponseAudit(audit)

	return c.Status(200).JSON(responseAudit)
}

func GetAudits(c *fiber.Ctx) error {
	audits := []models.Audit{}

	database.Database.Db.Find(&audits)
	reponseAudits := []Audit{}
	for _, audit := range audits {
		reponseAudit := CreateResponseAudit(audit)
		reponseAudits = append(reponseAudits, reponseAudit)
	}

	return c.Status(200).JSON(reponseAudits)
}

func findAudit(id int, audit *models.Audit) error {
	database.Database.Db.Find(&audit, "id = ?", id)
	if audit.ID == 0 {
		return errors.New("audit_role does not exist")
	}
	return nil
}

func GetAudit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var audit models.Audit

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findAudit(id, &audit); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseAudit := CreateResponseAudit(audit)

	return c.Status(200).JSON(responseAudit)

}

func DeleteAudit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var audit models.Audit

	if err != nil {
		return c.Status(400).JSON("ID is an integer")
	}

	if err := findAudit(id, &audit); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&audit).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted")

}
