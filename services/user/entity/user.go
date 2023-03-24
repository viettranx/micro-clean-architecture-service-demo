package entity

import (
	"demo-service/common"
	"github.com/viettranx/service-context/core"
	"strings"
)

type Gender string

const (
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
	GenderUnknown Gender = "unknown"
)

type SystemRole string

const (
	RoleSuperAdmin SystemRole = "sadmin"
	RoleAdmin      SystemRole = "admin"
	RoleUser       SystemRole = "user"
)

type Status string

const (
	StatusActive        Status = "active"
	StatusPendingVerify Status = "waiting_verify"
	StatusBanned        Status = "banned"
)

type User struct {
	core.SQLModel             // in practice, we should not embed this struct
	FirstName     string      `json:"first_name" gorm:"column:first_name" db:"first_name"`
	LastName      string      `json:"last_name" gorm:"column:last_name" db:"last_name"`
	Email         string      `json:"email" gorm:"column:email" db:"email"`
	Phone         string      `json:"phone" gorm:"column:phone" db:"phone"`
	Avatar        *core.Image `json:"avatar" gorm:"column:avatar" db:"avatar"`
	Gender        Gender      `json:"gender" gorm:"column:gender" db:"gender"`
	SystemRole    SystemRole  `json:"system_role" gorm:"column:system_role" db:"system_role"`
	Status        Status      `json:"status" gorm:"column:status" db:"status"`
}

func NewUser(firstName, lastName, email string) User {
	return User{
		SQLModel:   core.NewSQLModel(),
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Phone:      "",
		Avatar:     nil,
		Gender:     GenderUnknown,
		SystemRole: RoleUser,
		Status:     StatusActive,
	}
}

func (User) TableName() string { return "users" }

func (u *User) Validate() error {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(u.Email)

	if err := checkFirstName(u.FirstName); err != nil {
		return err
	}

	if err := checkLastName(u.LastName); err != nil {
		return err
	}

	if !emailIsValid(u.Email) {
		return ErrEmailIsNotValid
	}

	return nil
}

func (u *User) Mask() {
	u.SQLModel.Mask(common.MaskTypeUser)
}
