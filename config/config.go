package config

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	Host     = "localhost"
	User     = "yuji"
	Password = "S3cret"
	Name     = "payment"
	Port     = "5432"
)

func Setup() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		Host,
		Port,
		User,
		Name,
		Password,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
