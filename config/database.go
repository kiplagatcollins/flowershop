package config

import (
	"fmt"

	"gitlab.com/kiplagatcollins/flowershop/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "appuser"
	password = "secret"
	dbName   = "postgres"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sales.",
			SingularTable: true,
		},
	})
	db.AutoMigrate(model.User{}, model.Farmer{})

	if err != nil {
		panic("failed to connect to the Oracle database")
	}

	return db
}
