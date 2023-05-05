package user

import (
	"gestorpasswordapi/internal/utils"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type User struct {
	ID       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Salt     string `json:"salt,omitempty"`
}

func (u *User) HashPassword() error {
	// TODO: hash password in pbkdf2
	return nil
}

func (u *User) GenerateRandomSalt() error {
	var saltLengthStr = os.Getenv("SALT_LENGTH")
	saltLengthInt, err := strconv.Atoi(saltLengthStr)
	if err != nil {
		log.Fatal("Check SALT_LENGTH environment variable")
		return err
	}
	salt, err := utils.GenerateRandomString(saltLengthInt)
	if err != nil {
		log.Fatal("Error generating random salt")
		return err
	}
	u.Salt = salt
	return nil
}
