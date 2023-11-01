package utils

import (
	"io"
	"log"

	"CareerCourse/models"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
)

func Login(username, password string) error {
	strRSAPassword, strPassword, err := EncryPassWord(password)
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return err
	}
	payload := models.LoginPayload{
		LSchoolId:      0,
		StrLoginId:     username,
		StrRSAPassword: strRSAPassword,
		StrPassword:    strPassword,
		StrVerifyCode:  "",
	}

	res, err := Post("https://bistu.wnssedu.com/rest/v1/user/SSOLogin", payload.ToIOReader())
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return err
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return err
	}
	return nil
}
