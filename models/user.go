package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`
	ID            int64      `bun:",pk,autoincrement" json:"id"`
	Username      string     `json:"username"`
	Password      string     `json:"password"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Email         string     `json:"email"`
	UserTypeID    int64      `json:"user_type_id"`
	UserType      UserType   `bun:"rel:belongs-to,join:user_type_id=id" json:"user_type"`
	CreatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt     *time.Time `bun:",nullzero" json:"deleted_at"`
}

func GetAllUsers(db *bun.DB) ([]*User, error) {
	var users []*User
	user := User{}
	err := db.NewSelect().Model(&user).Relation("UserType").
		Scan(context.Background(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(db *bun.DB, id int64) (*User, error) {
	user := User{}
	err := db.NewSelect().Model(&user).Where("u.id = ?", id).
		Relation("UserType").Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return &user, nil
}
