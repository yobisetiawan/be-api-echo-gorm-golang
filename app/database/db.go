package database

import (
	"be_api/app/configs"
	"be_api/app/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var DB_CONNECT = false

// Initialize the database connection
func InitDB() {
	var err error
	// Setup PostgreSQL connection string
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		configs.AppConfig.DBUser,
		configs.AppConfig.DBPassword,
		configs.AppConfig.DBHost,
		configs.AppConfig.DBPort,
		configs.AppConfig.DBName,
		configs.AppConfig.DBSSLMode,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// Open the database connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("Unable to connect to DB", err)
	}

	DBAutoMigrate()
}

func DBAutoMigrate() {
	DB.AutoMigrate(
		models.Admin{},
		models.User{},
		models.Otp{},
		models.Session{},
		models.Product{},
		models.ProductCategory{},
	)
}
