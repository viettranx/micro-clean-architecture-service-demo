package entity

import (
	"github.com/viettranx/service-context/core"
	"strings"
)

// UserDataCreation use for inserting data into database, we don't need all data fields
type UserDataCreation struct {
	core.SQLModel
	FirstName string `json:"first_name" gorm:"column:first_name" db:"first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name" db:"last_name"`
	Email     string `json:"email" gorm:"column:email" db:"email"`
	// Do not allow client set these fields
	SystemRole SystemRole `json:"-" gorm:"column:system_role" db:"system_role"`
	Status     Status     `json:"-" gorm:"column:status" db:"status"`
}

func NewUserForCreation(firstName, lastName, email string) UserDataCreation {
	return UserDataCreation{
		SQLModel:   core.NewSQLModel(),
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		SystemRole: RoleUser,
		Status:     StatusActive,
	}
}

func (u *UserDataCreation) PrepareForInsert() {
	u.SQLModel = core.NewSQLModel()
	u.SystemRole = RoleUser
	u.Status = StatusActive
}

func (*UserDataCreation) TableName() string { return User{}.TableName() }

func (u *UserDataCreation) Validate() error {
	u.Email = strings.TrimSpace(u.Email)

	if ok := emailIsValid(u.Email); !ok {
		return ErrEmailIsNotValid
	}

	u.FirstName = strings.TrimSpace(u.FirstName)

	if err := checkFirstName(u.FirstName); err != nil {
		return err
	}

	u.LastName = strings.TrimSpace(u.LastName)

	if err := checkLastName(u.LastName); err != nil {
		return err
	}

	if err := checkStatus(u.Status); err != nil {
		return err
	}

	if err := checkRole(u.SystemRole); err != nil {
		return err
	}

	return nil
}

// UserDataUpdate contains only data fields can be used for updating
type UserDataUpdate struct {
	FirstName *string `json:"first_name" gorm:"column:first_name" db:"first_name"`
	LastName  *string `json:"last_name" gorm:"column:last_name" db:"last_name"`
	Phone     *string `json:"phone" gorm:"column:phone" db:"phone"`
	Gender    *Gender `json:"gender" gorm:"column:gender" db:"gender"`
	//Avatar    *core.Image `json:"avatar" gorm:"column:avatar" db:"avatar"`

	// Do not allow client set these fields
	SystemRole *SystemRole `json:"-" gorm:"column:system_role" db:"system_role"`
	Status     *Status     `json:"-" gorm:"column:status" db:"status"`
}

func (*UserDataUpdate) TableName() string { return User{}.TableName() }

func (u *UserDataUpdate) Validate() error {
	if firstName := u.FirstName; firstName != nil {
		s := strings.TrimSpace(*firstName)

		if err := checkFirstName(s); err != nil {
			return err
		}

		u.FirstName = &s
	}

	if lastName := u.LastName; lastName != nil {
		s := strings.TrimSpace(*lastName)

		if err := checkLastName(s); err != nil {
			return err
		}

		u.LastName = &s
	}

	if phone := u.Phone; phone != nil {
		s := strings.TrimSpace(*phone)

		if err := checkPhoneNumber(s); err != nil {
			return err
		}

		u.LastName = &s
	}

	if gender := u.Gender; gender != nil {
		if err := checkGender(*gender); err != nil {
			return err
		}
	}

	if status := u.Status; status != nil {
		if err := checkStatus(*status); err != nil {
			return err
		}
	}

	if role := u.SystemRole; role != nil {
		if err := checkRole(*role); err != nil {
			return err
		}
	}

	return nil
}
