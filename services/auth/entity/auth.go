package entity

import (
	"github.com/viettranx/service-context/core"
)

type Auth struct {
	core.SQLModel
	UserId     int    `json:"user_id" gorm:"column:user_id;" db:"user_id"`
	AuthType   string `json:"auth_type" gorm:"column:auth_type;" db:"auth_type"`
	Email      string `json:"email" gorm:"column:email;" db:"email"`
	Salt       string `json:"salt" gorm:"column:salt;" db:"salt"`
	Password   string `json:"password" gorm:"column:password;" db:"password"`
	FacebookId string `json:"facebook_id" gorm:"column:facebook_id" db:"facebook_id"`
}

func (Auth) TableName() string { return "auths" }

func NewAuthWithEmailPassword(userId int, email, salt, password string) Auth {
	return Auth{
		SQLModel: core.NewSQLModel(),
		UserId:   userId,
		Email:    email,
		Salt:     salt,
		Password: password,
		AuthType: "email_password",
	}
}
