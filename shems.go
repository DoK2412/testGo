package main

import (
	"time"
)

type Registr struct {
	Logins         string    `json:"logins"`
	Password       string    `json:"password"`
	RepeatPassword string    `json:"repeat_password"`
	DeviceInfo     DeviceInf `json:"device_info"`
	Test           string    `json:"test"`
}

type DeviceInf struct {
	PhoneNumber string `json:"phone_number"`
	Locale      string `json:"locale"`
}

type Сonfirm struct {
	Code int `json:"code"`
}

type Users struct {
	Id                int64     `gorm:"type: int64"`
	Logins            string    `gorm:"type: string"`
	Password          string    `gorm:"type: string"`
	Phone_number      string    `gorm:"type: string"`
	Locale            string    `gorm:"type: string"`
	Activated         bool      `gorm:"type: bool"`
	Registration_date time.Time `gorm:"type: time.Time"`
	Blocking_date     any       `gorm:"type: time.Time"`
	Code              int       `gorm:"type: int"`
}

type Authorization struct {
	Logins   string `json:"logins"`
	Password string `json:"password"`
}

type Setting struct {
	Locale string `json:"locale"`
}
