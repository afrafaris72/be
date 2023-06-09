package database

import (
	"dumbflix/models"
	"dumbflix/pkg/mysql"
	"fmt"
)

func RunMigrations() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Film{},
		&models.Category{},
		&models.Transaction{},
		&models.Episode{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration error")
	}

	fmt.Println("Migration succes")
}
