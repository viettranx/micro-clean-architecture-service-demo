package entity

import (
	"demo-service/common"
	"github.com/viettranx/service-context/core"
)

type Status string

const (
	StatusDoing   Status = "doing"
	StatusDone    Status = "done"
	StatusDeleted Status = "deleted"
)

type Task struct {
	core.SQLModel
	UserId      int              `json:"-" gorm:"column:user_id" db:"user_id"`
	Title       string           `json:"title" gorm:"column:title;" db:"title"`
	Description string           `json:"description" gorm:"column:description;" db:"description"`
	Status      Status           `json:"status" gorm:"column:status;" db:"status"`
	User        *core.SimpleUser `json:"user" gorm:"-" db:"-"`
}

func (Task) TableName() string { return "tasks" }

func (t *Task) Mask() {
	t.SQLModel.Mask(common.MaskTypeTask)

	if u := t.User; u != nil {
		u.Mask(common.MaskTypeUser)
	}
}
