package controllers

import (

	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/Vafa/inputs"
	"github.com/mhd7966/Vafa/log"
	"github.com/mhd7966/Vafa/models"
	"github.com/mhd7966/Vafa/repositories"
	"github.com/sirupsen/logrus"
)

// GetUsers godoc
// @Summary get users
// @Description return users
// @ID get_users
// @Tags User
// @Param page_status query inputs.GetUsersQuery true "page and query"
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /users/ [get]
func GetUsers(c *fiber.Ctx) error {
	var response models.Response

	var query inputs.GetUsersQuery
	if err := c.QueryParser(&query); err != nil {
		response.Message = "Get Query Body Failed!"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("GetَUsers. Get users query body failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	users, err := repositories.GetUsers(&query)
	if err != nil {
		response.Message = "Get users Failed!"
		log.Log.WithFields(logrus.Fields{
			"users":      users,
			"query_body": query,
			"response":   response.Message,
			"error":      err.Error(),
		}).Error("GetUsers. Get users from DB failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Message = "OK!"
	response.Status = "succes"
	response.Data = users
	log.Log.WithFields(logrus.Fields{
		"users":      users,
		"query_body": query,
		"response":   response.Message,
	}).Info("GetUsers. Get users successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}
