package models

import (
	"context"

	"github.com/uptrace/bun"
)

type UserType struct {
	bun.BaseModel `bun:"table:user_type"`
	ID            int64  `json:"id"`
	Name          string `json:"name"`
}

func GetAllUserTypes(db *bun.DB) ([]*UserType, error) {
	var userTypes []*UserType
	user := UserType{}
	ctx := context.Background()
	err := db.NewSelect().Model(&user).Scan(ctx, &userTypes)
	if err != nil {
		return nil, err
	}
	return userTypes, nil
}
