package helpers

import (
	"fmt"
	"os"
	"qayyuum/go_fintech/interfaces"
	"regexp"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	HOST    = "127.0.0.1"
	DB_PORT = "5432"
	DB_NAME = "fintech"
)

func HashAndSalt(pass []byte) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func ConnectDB() (*gorm.DB, error) {
	db_user := os.Getenv("DB_USER")
	if db_user == "" {
		return nil, fmt.Errorf("DB_USER is not provided")
	}

	db_password := os.Getenv("DB_PASSWORD")
	if db_password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is not provided")
	}

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", HOST, DB_PORT, db_user, DB_NAME, db_password))
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}
	return db, nil
}

func Validation(values []interfaces.Validation) bool {
	username := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$`)

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}
	return true
}
