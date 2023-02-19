package controller

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

/*
	Util functions is used for supporting main API implementation functions, and not exposed publicly
*/

func (c *accountServiceController) generateRandomPassword() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	res := make([]byte, 10) // generate random password fixed in 10-char size
	for i := range res {
		res[i] = letters[rand.Intn(52)] // 52 is total number of english alphabet letters
	}
	return string(res)
}

func (c *accountServiceController) validatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}

func (c *accountServiceController) hashPasswordBcrypt(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func (c *accountServiceController) checkPasswordBcrypt(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
