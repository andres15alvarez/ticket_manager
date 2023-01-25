package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Ticket struct {
	bun.BaseModel       `bun:"table:ticket"`
	ID                  int64 `bun:",pk,autoincrement"`
	CustomerName        string
	CustomerEmail       string
	CustomerPhoneNumber string
	InvoiceNumber       string
	WarantyNumber       string
	Detail              string
	InvoiceImage        string
	WarantyCertificate  string
	Subject             string
	Answer              string
	HelpTopicID         int64
	HelpTopic           *HelpTopic `bun:"rel:belongs-to,join:help_topic_id=id"`
	DepartmentId        int64
	Department          *Department `bun:"rel:belongs-to,join:department_id=id"`
	StateID             int64
	State               *State `bun:"rel:belongs-to,join:state_id=id"`
	CreatedByID         int64
	CreatedBy           *User `bun:"rel:belongs-to,join:created_by_id=id"`
	TakenByID           int64
	TakenBy             *User     `bun:"rel:belongs-to,join:taken_by_id=id"`
	TakenAt             time.Time `bun:",nullzero"`
	CreatedAt           time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt           time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt           time.Time `bun:",nullzero"`
}
