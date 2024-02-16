package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:varchar(36)"`
	Name     string    `json:"name" gorm:"type:text"`
	Surname  string    `json:"surname" gorm:"type:text"`
	Username string    `json:"username" gorm:"type:varchar(255);unique"`
	Email    string    `json:"email" gorm:"type:varchar(255);unique"`
	Password string    `json:"password" gorm:"type:text"`
}
