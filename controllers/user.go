package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/Vafa/inputs"
	"github.com/mhd7966/Vafa/log"
	"github.com/mhd7966/Vafa/models"
	"github.com/mhd7966/Vafa/repositories"
	"github.com/sirupsen/logrus"
)

// GetUser godoc
// @Summary get user
// @Description return user
// @ID get_user
// @Tags user
// @Param id path int true "user_id"
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	var response models.Response
	response.Status = "error"

	userID, err := c.ParamsInt("id")
	if err != nil {
		response.Message = "Convert Param Failed!"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("GetUser. Convert param string to int failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if existUser := repositories.ExistUserByID(uint(userID)); !existUser {
		response.Message = "User Not Found!"
		log.Log.WithFields(logrus.Fields{
			"user_id":  userID,
			"response": response.Message,
		}).Error("DeleteUser. User doesn't exist!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user, err := repositories.GetUser(userID)
	if err != nil {
		response.Message = "Get user Failed!"
		log.Log.WithFields(logrus.Fields{
			"user_id":  userID,
			"response": response.Message,
			"error":    err.Error(),
		}).Error("GetUser. Get user from DB failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Message = "OK!"
	response.Status = "succes"
	response.Data = user
	log.Log.WithFields(logrus.Fields{
		"user":     user,
		"response": response.Message,
	}).Info("GetUser. Get user successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}

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
	response.Status = "error"

	var query inputs.GetUsersQuery
	if err := c.QueryParser(&query); err != nil {
		response.Message = "Get Query Body Failed!"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("GetَUsers. Get users query body failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	users, err := repositories.GetUsers(query)
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

// NewUser godoc
// @Summary new user
// @Description create user
// @ID new_user
// @Tags user
// @Param userBody body models.User true "user"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /users [post]
func NewUser(c *fiber.Ctx) error {
	var response models.Response
	response.Status = "error"
	userBody := new(models.User)

	if err := c.BodyParser(userBody); err != nil {
		response.Message = "Parse Body Failed"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("NewUser. Parse body to user body failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	validate := validator.New()
	if err := validate.Struct(userBody); err != nil {
		response.Message = "Validate NationalID Failed"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("NewUser. Validate user info failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if existUser := repositories.ExistUserByNationalID(userBody.NationalID); existUser {
		response.Message = "Duplicate User"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
		}).Error("NewUser. This nationalID is duplicate. we have one of this!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := repositories.CreateUser(userBody); err != nil {
		response.Message = "Create User Failed"
		log.Log.WithFields(logrus.Fields{
			"user":     userBody,
			"response": response.Message,
			"error":    err.Error(),
		}).Error("NewUser. Create user Failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Message = "OK!"
	response.Status = "succes"
	response.Data = userBody
	log.Log.WithFields(logrus.Fields{
		"user":     userBody,
		"response": response.Message,
	}).Info("NewUser. Create User successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}

// CancelUser godoc
// @Summary cancel user
// @Description cancel user
// @ID cancel_user
// @Tags user
// @Param id path int true "user_id"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /users/{id}/cancel [put]
func CancelUser(c *fiber.Ctx) error {
	var response models.Response
	response.Status = "error"

	userID, err := c.ParamsInt("id")
	if err != nil {
		response.Message = "Convert Param Failed!"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("CancelUser. Convert param string to int failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if existUser := repositories.ExistUserByID(uint(userID)); !existUser {
		response.Message = "User Not Found!"
		log.Log.WithFields(logrus.Fields{
			"user_id":  userID,
			"response": response.Message,
		}).Error("CancelUser. User doesn't exist!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err = repositories.CancelUser(userID); err != nil {
		response.Message = "Cancel User Failed"
		log.Log.WithFields(logrus.Fields{
			"userID":   userID,
			"response": response.Message,
			"error":    err.Error(),
		}).Error("CancelUser. Cancel user in DB failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Message = "OK!"
	response.Status = "succes"
	log.Log.WithFields(logrus.Fields{
		"response": response.Message,
	}).Info("CancelUser. Cancel user successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}

// DeleteUser godoc
// @Summary delete user
// @Description delete user
// @ID delete_user
// @Tags user
// @Param id path int true "user_id"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	var response models.Response
	response.Status = "error"

	userID, err := c.ParamsInt("id")
	if err != nil {
		response.Message = "Convert Param Failed!"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("DeleteUser. Convert param string to int failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if existUser := repositories.ExistUserByID(uint(userID)); !existUser {
		response.Message = "User Not Found!"
		log.Log.WithFields(logrus.Fields{
			"user_id":  userID,
			"response": response.Message,
		}).Error("DeleteUser. User doesn't exist!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err = repositories.DeleteUser(userID); err != nil {
		response.Message = "Delete User Failed!"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("DeleteUser. delete user from DB failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Message = "OK!"
	response.Status = "succes"
	log.Log.WithFields(logrus.Fields{
		"response": response.Message,
	}).Info("DeleteUser. Delete user successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}
