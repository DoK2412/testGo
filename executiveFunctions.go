package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"goTestAPI/database"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"regexp"
	"time"
)

func HashPassword(password string) (string, error) {
	//Функция хеширования пароля
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Registration(content Registr, session sessions.Session) map[string]string {
	// Функция регистрации пользователя
	connector, err := database.Connection()
	if err != nil {
		return APIAnswer("90")
	}

	var addUser Users
	connector.First(&addUser, fmt.Sprintf("logins = '%s'", content.Logins)).First(&addUser)
	if addUser.Id != 0 {
		return APIAnswer("13")
	}

	if content.Password != content.RepeatPassword {
		return APIAnswer("10")
	}

	checkNumber, err := regexp.MatchString("^((8|\\+7)[\\- ]?)?(\\(?\\d{3}\\)?[\\- ]?)?[\\d\\- ]{10,11}$", content.DeviceInfo.PhoneNumber)
	if err != nil {
		log.Fatal(err)
	} else if !checkNumber {
		return APIAnswer("12")
	}

	var code int
	if content.Test != "dev" {
		code = rand.Intn(10000) + 10000
	} else {
		code = 6666
	}
	hash, _ := HashPassword(content.Password)

	saveUser := []*Users{
		{Logins: content.Logins,
			Password:          hash,
			Phone_number:      content.DeviceInfo.PhoneNumber,
			Locale:            content.DeviceInfo.Locale,
			Activated:         false,
			Registration_date: time.Now(),
			Code:              code},
	}

	connector.Create(saveUser)

	session.Set("userName", content.Logins)
	session.Save()

	return APIAnswer("0")

}

func Confirmation(content Сonfirm, session sessions.Session) map[string]string {
	// Функция подтверждения учетной записи пользователя
	connector, err := database.Connection()
	if err != nil {
		return APIAnswer("91")
	}
	userLogin := session.Get("userName")
	var addUser Users

	connector.First(&addUser, fmt.Sprintf("logins = '%s'", userLogin)).First(&addUser)
	if addUser.Code != content.Code {
		return APIAnswer("14")
	} else {
		session.Set("userID", addUser.Id)
	}

	connector.Model(addUser).Where(fmt.Sprintf("logins = '%s'", userLogin)).Update("activated", true)

	return APIAnswer("0")

}

func Authorizations(content Authorization, session sessions.Session) map[string]string {
	connector, err := database.Connection()
	if err != nil {
		return APIAnswer("91")
	}

	var addUser Users
	connector.First(&addUser, fmt.Sprintf("logins = '%s'", content.Logins)).First(&addUser)
	if CheckPassword(content.Password, addUser.Password) {
		session.Set("userID", addUser.Id)
		session.Set("userName", addUser.Logins)
		session.Set("userPhoneNumber", addUser.Phone_number)
		session.Set("userLocale", addUser.Locale)
		session.Save()
		return APIAnswer("0")

	} else {
		return APIAnswer("15")
	}

}

func Settings(content Setting, session sessions.Session) map[string]string {

	if session.Get("userID") != nil {
		connector, err := database.Connection()
		if err != nil {
			return APIAnswer("91")
		}
		userLogin := session.Get("userName")
		var addUser Users
		connector.Model(addUser).Where(fmt.Sprintf("logins = '%s'", userLogin)).Update("locale", content.Locale)
		return APIAnswer("0")
	} else {
		return APIAnswer("20")
	}

}
