package database

import "github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/models"

func Migrate() {

	var err error

	if err = DB.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

}
