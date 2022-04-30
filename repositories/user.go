package repositories

import (
	"github.com/mhd7966/Vafa/connections"
	"github.com/mhd7966/Vafa/inputs"
	"github.com/mhd7966/Vafa/models"
)

func ExistUser(NationalID string) bool {
	var user models.User

	if result := connections.DB.Where(models.User{NationalID: NationalID}).First(&user); result.RowsAffected == 0 {
		return false
	}
	return true
}

func GetUsers(query *inputs.GetUsersQuery) (*[]models.User, error) {
	var users []models.User

	offset := (query.Page - 1) * query.PageSize
	if result := connections.DB.Where("status = ?", query.Status).Offset(offset).Limit(query.PageSize).Find(&users); result.RowsAffected < 1 {
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
