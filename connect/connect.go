package connect

import (
	"fmt"

	"github.com/testfiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	dsn := "root:root@/go_admin"
	v := "Cannnot find the connection"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(v)
	}

	fmt.Println(db)

	Database = db

	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Product{}, &models.Category{}, &models.Picture{}, &models.Cart{})

}
