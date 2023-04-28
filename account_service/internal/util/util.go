package util

import (
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

/*
	Util functions is used for supporting main API implementation functions, and not exposed publicly
*/

func GenerateRandomPassword() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	res := make([]byte, 6) // generate random password fixed in 10-char size
	for i := range res {
		res[i] = letters[rand.Intn(52)] // 52 is total number of english alphabet letters
	}
	return string(res)
}

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}

func HashPasswordBcrypt(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func CheckPasswordBcrypt(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func SendEmail(to []string, title string, message string) error {
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	smtpUsername := "gaokotrang@gmail.com"
	smtpPassword := "ugresofpeawzmuds"

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)

	// convert message to byte][]
	smg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + title + "\r\n" +
		"\r\n" +
		message + "\r\n")

	// send email
	return smtp.SendMail(smtpServer+":"+string(smtpPort), auth, smtpUsername, to, smg)

}

func GenerateOTPCode() string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000000)
	return fmt.Sprintf("%06d", randomNum)
}
