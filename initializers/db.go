package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"qdjr-api/helpers"
	"qdjr-api/models"
)

type DbInitializer struct{}

var DB *gorm.DB

var baseHelper = new(helpers.BaseHelper)

func (_ DbInitializer) Database() *gorm.DB {
	return DB
}

func (dbInitializer DbInitializer) ConnectDataBase() {
	var err error

	DBDriver := baseHelper.GetEnv("DB_DRIVER", "postgres")
	DBUser := baseHelper.GetEnv("DB_USER", "postgres_user")
	DBPassword := baseHelper.GetEnv("DB_PASSWORD", "12345@")
	DBName := baseHelper.GetEnv("DB_NAME", "postgres")
	DBHost := baseHelper.GetEnv("DB_HOST", "localhost")
	DBPort := baseHelper.GetEnv("DB_PORT", "5432")

	DBUri := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBHost,
		DBUser,
		DBPassword,
		DBName,
		DBPort,
	)
	switch DBDriver {
	case "postgres":
		DB, err = gorm.Open(postgres.Open(DBUri))
		break
	default:
		err = fmt.Errorf("unknown driver %s", DBDriver)
	}

	if err != nil {
		fmt.Println("Cannot connect to database ", DBDriver)
		log.Fatal("Connection error: ", err)
	} else {
		fmt.Println("We are connected to the database ", DBDriver)
	}
	dbInitializer.autoMigrate()
}

func (_ DbInitializer) autoMigrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Article{})
	if err != nil {
		fmt.Println("Cannot auto migrate database with err: ", err)
		log.Fatal("Cannot auto migrate database with err: ", err)
	} else {
		fmt.Println("Migrate success")
	}
}
