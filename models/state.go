package models

import (
	"context"

	"github.com/uptrace/bun"
)

type State struct {
	bun.BaseModel `bun:"table:state"`
	ID            int64  `bun:",pk,autoincrement" json:"id"`
	Name          string `json:"name"`
}

func GetAllStates(db *bun.DB) ([]*State, error) {
	var states []*State
	state := State{}
	err := db.NewSelect().Model(&state).Scan(context.Background(), &states)
	if err != nil {
		return nil, err
	}
	return states, nil
}
