package entity

import (
	"time"

	"github.com/viettranx/service-context/core"
)

type Status string

const (
	// StatusDoing indicates task is in progress.
	StatusDoing Status = "doing"
	// StatusDone indicates task is completed.
	StatusDone Status = "done"
	// StatusDeleted indicates task is deleted.
	StatusDeleted Status = "deleted"
)

// Task defines data model for task.
type Task struct {
	ID          uint        `json:"-" gorm:"primary_key;column:id;auto_increment:true"`
	FakeID      *core.UID   `json:"id" gorm:"-"`
	UserID      uint        `json:"user_id" gorm:"column:user_id"`
	User        *SimpleUser `json:"user" gorm:"-"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      Status      `json:"status" gorm:"column:status"`
	CreatedAt   time.Time   `json:"created_at" gorm:"column:created_at;autocreatetime"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"column:updated_at;autoupdatetime"`
}

func (Task) TableName() string { return "tasks" }

// SimpleUser only contains public infos.
type SimpleUser struct {
	ID        uint      `json:"-"`
	FakeID    *core.UID `json:"id"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
}

const (
	MaskTypeUser = iota + 1
	MaskTypeTask
)

func (u *SimpleUser) Mask() {
	uid := core.NewUID(uint32(u.ID), MaskTypeUser, 1)
	u.FakeID = &uid
}

func (t *Task) Mask() {
	uid := core.NewUID(uint32(t.ID), MaskTypeTask, 1)
	t.FakeID = &uid

	if u := t.User; u != nil {
		u.Mask()
	}
}
