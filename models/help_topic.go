package models

import "github.com/uptrace/bun"

type HelpTopic struct {
	bun.BaseModel `bun:"table:help_topic"`
	ID            int64
	Name          string
}
