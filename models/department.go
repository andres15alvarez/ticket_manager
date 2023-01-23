package models

import "github.com/uptrace/bun"

type Department struct {
	bun.BaseModel `bun:"table:department"`
	ID            int64
	Name          string
}
