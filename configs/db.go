package configs

import (
	"fmt"
	"os"

	"github.com/dezenter/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a shared connection to the database
var DB *gorm.DB

// InitDB creates a shared connection to the database
func InitDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbn := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USERNAME")
	pwd := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable TimeZone=Asia/Bangkok", host, port, dbn, user, pwd)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	envApp := os.Getenv("APP_ENV")
	if envApp == "dev" {
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.PostCategory{})
		db.AutoMigrate(&models.Post{})
		db.AutoMigrate(&models.Admin{})
		// db.AutoMigrate(&models.ResetPassword{})
		// db.AutoMigrate(&models.Event{})
	}

	DB = db
}
