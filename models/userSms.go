// Package models provides 
package models

type UserSms struct {
	Id        int
	Ip        string
	Phone     string
	SendCount int
	AddDay    string
	AddTime   int
	Sign      string
}

func (u *UserSms) TableName() string {
	return "user_sms"
}
