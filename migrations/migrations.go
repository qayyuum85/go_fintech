package migrations

import (
	"fmt"
	"qayyuum/go_fintech/helpers"
	"qayyuum/go_fintech/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() error {
	db, err := helpers.ConnectDB()
	if err != nil {
		return fmt.Errorf("Error connecting to database: %v", err)
	}

	users := &[2]interfaces.User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword, err := helpers.HashAndSalt([]byte(users[i].Username))
		if err != nil {
			fmt.Printf("Unable to generate password for user: %s", users[i].Username)
			continue
		}
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()

	return nil
}

// Migrate - run migration
func Migrate() error {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db, err := helpers.ConnectDB()
	if err != nil {
		return err
	}
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccounts()

	return nil
}
