package migrations

import (
	"qayyuum/go_fintech/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User model
type User struct {
	gorm.Model
	Username string
	Email string
	Password string
}

// Account model
type Account struct {
	gorm.Model
	Type string
	Name string
	Balance uint
	UserID uint
}

func connectDB() *gorm.DB {
    db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=fintech password=demo1234 sslmode=disable")
    helpers.HandleErr(err)
    return db
}


func createAccounts() {
    db := connectDB()

    users := [2]User{
        {Username: "Martin", Email: "martin@martin.com"},
        {Username: "Michael", Email: "michael@michael.com"},
    }

    for i := 0; i < len(users); i++ {
        generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
        user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
        db.Create(&user)

        account := Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
        db.Create(&account)
    }
    defer db.Close()
}

// Migrate - run migration
func Migrate() {
    db := connectDB()
    db.AutoMigrate(&User{}, &Account{})
    defer db.Close()

    createAccounts()
}