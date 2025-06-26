package helpers

import (
	"ewallet-ums/internal/models"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetUpMySQL() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnv("DB_USER", "root"),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_PORT", "3306"),
		GetEnv("DB_NAME", "ewallet_db"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to MySQL: " + err.Error())
	}

	logrus.Info("Connected to MySQL database successfully")
	DB.AutoMigrate(
		&models.User{},
		&models.UserSession{},
	)
}
