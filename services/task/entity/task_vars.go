package entity

import (
	"demo-service/common"
	"github.com/viettranx/service-context/core"
	"strings"
)

// TaskDataCreation use for inserting data into database, we don't need all data fields
type TaskDataCreation struct {
	core.SQLModel
	Title       string `json:"title" gorm:"column:title;" db:"title"`
	Description string `json:"description" gorm:"column:description;" db:"description"`
	// Do not allow client set these fields
	UserId int    `json:"-" gorm:"column:user_id" db:"user_id"`
	Status Status `json:"-" gorm:"column:status;" db:"status"`
}

func (TaskDataCreation) TableName() string { return Task{}.TableName() }

func (t *TaskDataCreation) Prepare(userId int, status Status) {
	t.SQLModel = core.NewSQLModel()
	t.UserId = userId
	t.Status = status
}

func (t *TaskDataCreation) Mask() {
	t.SQLModel.Mask(common.MaskTypeTask)
}

func (t *TaskDataCreation) Validate() error {
	t.Title = strings.TrimSpace(t.Title)

	if err := checkTitle(t.Title); err != nil {
		return err
	}

	if err := checkStatus(t.Status); err != nil {
		return err
	}

	if t.UserId <= 0 {
		return ErrUserIdNotValid
	}

	return nil
}

// TaskDataUpdate contains only data fields can be used for updating
type TaskDataUpdate struct {
	Title       *string `json:"title" gorm:"column:title;" db:"title"`
	Description *string `json:"description" gorm:"column:description;" db:"description"`
	// Do not allow client set these fields
	Status *Status `json:"-" gorm:"column:status;" db:"status"`
}

func (TaskDataUpdate) TableName() string { return Task{}.TableName() }

func (t *TaskDataUpdate) Validate() error {
	if title := t.Title; title != nil {
		s := strings.TrimSpace(*title)

		if err := checkTitle(s); err != nil {
			return err
		}

		t.Title = &s
	}

	if status := t.Status; status != nil {
		if err := checkStatus(*status); err != nil {
			return err
		}
	}

	return nil
}

type Filter struct {
	UserId *string `json:"user_id,omitempty" form:"user_id"`
	Status *string `json:"status,omitempty" form:"status"`
}
