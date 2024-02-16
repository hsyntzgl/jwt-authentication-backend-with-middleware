package repository

import (
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/models"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/pkg/database"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/pkg/hasher"
)

func CrateUser(u *models.User) error {
	var err error
	db := database.DB

	if err = db.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string, password string) (*models.User, error) {
	var err error
	db := database.DB

	var user models.User

	if err = db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err = checkUserPassword(password, user); err != nil {
		return nil, err
	}
	return &user, nil
}
func GetUserByUsername(username string, password string) (*models.User, error) {
	var err error
	db := database.DB

	var user models.User

	if err = db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	if err = checkUserPassword(password, user); err != nil {
		return nil, err
	}

	return &user, nil

}
func checkUserPassword(password string, user models.User) error {
	if err := hasher.CheckPasswordHash(password, user.Password); err != nil {
		return err
	}
	return nil
}

func GetUserByID(uuid string) error {
	var err error
	db := database.DB

	var user models.User

	if err = db.Model(&models.User{}).First(&user).Where("uuid = ?", uuid).Error; err != nil {
		return err
	}

	return nil
}
