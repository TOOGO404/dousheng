package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GetPwd(pwd string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(b)
}

func ComparePwd(hash string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	log.Println(err)
	return err == nil
}
