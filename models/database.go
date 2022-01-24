package models

import (
	"fmt"

	"github.com/adijha/omicorn/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the he database connection.
var DB *gorm.DB

// SetupDatabase migrates and sets up the database.
func SetupDatabase() {
	u := helper.GetEnv("DATABASE_USER", "aditya.jha")
	p := helper.GetEnv("DATABASE_PASSWORD", "password")
	h := helper.GetEnv("DATABASE_HOST", "localhost:5432")
	n := helper.GetEnv("DATABASE_NAME", "omicorn")
	q := "charset=utf8mb4&parseTime=True&loc=Local"

	// Assemble the connection string.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", u, p, h, n, q)

	// Connect to the database.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(&User{})

	if err != nil {
		panic("Could not open database connection")
	}

	DB = db
}
