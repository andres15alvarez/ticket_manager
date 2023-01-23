package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`
	ID            int64
	Username      string
	Password      string
	FirstName     string
	LastName      string
	Email         string
	UserTypeID    int64
	UserType      *UserType `bun:"rel:belongs-to,join:user_type_id:id"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `bun:",nullzero"`
}
