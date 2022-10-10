package database

import (
	"fmt"
	"gofiber-penitipan-barang/application/config"
	"gofiber-penitipan-barang/application/model"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("port error")
	}
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	var dsn string

	if config.Config("DB_DRIVER") == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if config.Config("DB_DRIVER") == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	fmt.Println(dsn)

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := DB.DB()

	if err != nil {
		panic("error db in 42 line")
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{}, &model.Category{})
	fmt.Println("Database Migrated")
}
