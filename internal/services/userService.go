package services

import (
	"github.com/google/uuid"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/internal/repository"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/models"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/pkg/hasher"
)

func CreateUser(u *models.User) error {
	var err error

	u.ID = uuid.New()

	if u.Password, err = hasher.HashPassword(u.Password); err != nil {
		return err
	}

	err = repository.CrateUser(u)
	if err != nil {
		return err
	}

	return nil
}
