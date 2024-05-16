package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"haha/internal/logger"
	"haha/internal/models"
)

var DB *gorm.DB

func InitializeDB(host, port, user, password, name string, logg logger.Logger) (*gorm.DB, error) {
	var err error
	urlPostgres := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)

	DB, err = gorm.Open("postgres", urlPostgres)
	if err != nil {
		logg.Panic(err)
		return nil, err
	}

	DB.AutoMigrate(&models.User{}, &models.Vacancy{}, &models.Resume{},
		&models.Work{}, &models.ResToVac{}, &models.VacToRes{})

	err = DB.DB().Ping()
	if err != nil {
		logg.Panic(err)
		return nil, err
	}
	logg.Info("init database")

	return DB, nil
}

func GetDB() *gorm.DB {
	return DB
}
