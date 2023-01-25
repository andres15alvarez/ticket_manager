package models

import (
	"context"

	"github.com/uptrace/bun"
)

type Department struct {
	bun.BaseModel `bun:"table:department"`
	ID            int64  `bun:",pk,autoincrement" json:"id"`
	Name          string `json:"name"`
}

func GetAllDepartments(db *bun.DB) ([]*Department, error) {
	var departments []*Department
	department := Department{}
	ctx := context.Background()
	err := db.NewSelect().Model(&department).Scan(ctx, &departments)
	if err != nil {
		return nil, err
	}
	return departments, nil
}
