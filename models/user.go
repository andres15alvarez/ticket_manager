package models

import (
	"context"
	"log"
	"time"

	"github.com/andres15alvarez/ticket_manager/utils"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`
	ID            int64      `bun:",pk,autoincrement" json:"id"`
	Username      string     `bun:",unique" json:"username"`
	Password      string     `json:"-"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Email         string     `bun:",unique" json:"email"`
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

func GetUserByEmail(db *bun.DB, email string) (*User, error) {
	user := User{}
	err := db.NewSelect().Model(&user).Where("email = ?", email).
		Relation("UserType").Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(db *bun.DB, user *User) (*User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.UserTypeID = utils.UserTypeID
	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	res, err := db.NewInsert().Model(user).Exec(context.Background())
	log.Printf("User Created: %v", res)
	if err != nil {
		return nil, err
	}
	userCreated, err := GetUserByEmail(db, user.Email)
	if userCreated == nil {
		return user, nil
	}
	return userCreated, nil
}
