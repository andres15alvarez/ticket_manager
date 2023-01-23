package models

import "github.com/uptrace/bun"

type State struct {
	bun.BaseModel `bun:"table:state"`
	ID            int64
	Name          string
}
