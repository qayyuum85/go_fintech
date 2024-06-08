package users

import (
	"fmt"
	"qayyuum/go_fintech/helpers"
	"qayyuum/go_fintech/interfaces"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Login - login function signed with jwt token
func Login(username string, pass string) (map[string]interface{}, error) {
	isValidUser := helpers.Validation([]interfaces.Validation{
		{Valid: "username", Value: username},
		{Valid: "password", Value: pass},
	})
	if !isValidUser {
		return map[string]interface{}{"message": "invalid user"}, nil
	}

	db, err := helpers.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Check for username
	user := &interfaces.User{}
	if db.Where("username = ? ", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}, nil
	}

	// Compare the password with db
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong password"}, nil
	}

	accounts := []interfaces.ResponseAccount{}

	// get account info from db
	db.Table("accounts").Select("id, name, balance").Where("user_id = ? ", user.ID).Scan(&accounts)

	return prepareResponse(user, accounts)
}

func prepareToken(user *interfaces.User) (string, error) {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}

	// sign the connection with jwt token
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount) (map[string]interface{}, error) {
	// set the response user
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	token, err := prepareToken(user)
	if err != nil {
		return nil, fmt.Errorf("Unable to prepare token %v", err)
	}

	var response = map[string]interface{}{"message": "Login successful"}
	response["jwt"] = token
	response["data"] = responseUser

	return response, nil
}

func Register(username, email, pass string) error {
	return nil
}
