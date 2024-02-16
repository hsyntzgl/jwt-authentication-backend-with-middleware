package database

import (
	"fmt"
	"strconv"

	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"), "?parseTime=true")

	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}
}
