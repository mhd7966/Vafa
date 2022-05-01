package repositories

import (

	"github.com/mhd7966/Vafa/connections"
	"github.com/mhd7966/Vafa/constants"
	"github.com/mhd7966/Vafa/inputs"
	"github.com/mhd7966/Vafa/models"
)

func ExistUserByID(userID uint) bool {
	var user models.User

	if result := connections.DB.First(&user, userID); result.RowsAffected == 0 {
		return false
	}
	return true
}

func ExistUserByNationalID(nationalID string) bool {
	var user models.User

	if result := connections.DB.Where(models.User{
		NationalID: nationalID,
	}).First(&user); result.RowsAffected == 0 {
		return false
	}
	return true
}

func GetUser(userID int) (*models.User, error) {
	var user models.User

	if result := connections.DB.First(&user, userID); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUsers(query inputs.GetUsersQuery) (*[]models.User, error) {
	var users []models.User

	offset := (query.Page - 1) * query.Size

	if result := connections.DB.Where("status = ?", query.Status).Offset(offset).Limit(query.Size).Find(&users); result.RowsAffected < 1 {
		return nil, result.Error
	}

	return &users, nil
}

func CreateUser(user *models.User) error {

	if result := connections.DB.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func CancelUser(userID int) error {

	if result := connections.DB.Model(&models.User{}).Where("id = ?", userID).Update("status", constants.CANCEL); result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(userID int) error {

	if result := connections.DB.Where("id = ?", userID).Delete(&models.User{}); result.Error != nil {
		return result.Error
	}

	return nil
}
