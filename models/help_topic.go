package models

import (
	"context"

	"github.com/uptrace/bun"
)

type HelpTopic struct {
	bun.BaseModel `bun:"table:help_topic"`
	ID            int64  `bun:",pk,autoincrement" json:"id"`
	Name          string `json:"name"`
}

func GetAllHelpTopics(db *bun.DB) ([]*HelpTopic, error) {
	var helpTopics []*HelpTopic
	helpTopic := HelpTopic{}
	ctx := context.Background()
	err := db.NewSelect().Model(&helpTopic).Scan(ctx, &helpTopics)
	if err != nil {
		return nil, err
	}
	return helpTopics, nil
}
