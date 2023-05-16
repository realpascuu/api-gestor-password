package user

import (
	"crypto/sha512"
	"gestorpasswordapi/internal/utils"
	"log"
	"os"
	"strconv"

	"encoding/base64"

	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/pbkdf2"
)

type User struct {
	ID       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Salt     string `json:"salt,omitempty"`
}

func (u *User) GeneratePasswordHash(password string) string {
	iter, err := strconv.Atoi(os.Getenv("ITER_PASSWD"))
	if err != nil {
		return ""
	}
	keyLen, err := strconv.Atoi(os.Getenv("KEYLEN_PASSWD"))
	if err != nil {
		return ""
	}
	hashedPassword := pbkdf2.Key([]byte(password), []byte(u.Salt), iter, keyLen, sha512.New)
	return base64.StdEncoding.EncodeToString(hashedPassword)
}

func (u *User) PasswordMatch(password string) bool {
	hashPassword := u.GeneratePasswordHash(password)
	return u.Password == hashPassword
}

func (u *User) GenerateRandomSalt() (string, error) {
	var saltLengthStr = os.Getenv("SALT_LENGTH")
	saltLengthInt, err := strconv.Atoi(saltLengthStr)
	if err != nil {
		log.Fatal("Check SALT_LENGTH environment variable")
		return "", err
	}
	salt, err := utils.GenerateRandomString(saltLengthInt)
	if err != nil {
		log.Println("Error generating random salt")
		return "", err
	}
	return salt, nil
}
