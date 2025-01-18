package entity

import (
	"time"
)

type Status string

const (
	StatusNew      Status = "new"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Order struct {
	ID        uint64
	UUID      int
	Status    Status
	Src       string
	Dst       string
	User      *User
	CreatedAt time.Time
	UpdateAt  *time.Time
}
