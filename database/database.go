package database

import (
	"fmt"
	"go-twitter/entities"
	"go-twitter/utils"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

func init() {
	err := godotenv.Load()

	if err != nil {
		utils.WarningLog.Println("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v",
		os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"),
		os.Getenv("PG_DBNM"), os.Getenv("PG_PORT"),
	)

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	}

	utils.InfoLog.Println("connecting to database")
	Conn, err = gorm.Open(postgres.Open(dsn), gormConfig)
	Conn.AutoMigrate(&entities.User{})
	if err != nil {
		utils.ErrorLog.Panicln(err.Error())
	}
}
