package models

import (
	"context"
	"time"

	"github.com/andres15alvarez/ticket_manager/utils"

	"github.com/uptrace/bun"
)

type Ticket struct {
	bun.BaseModel       `bun:"table:ticket"`
	ID                  int64       `bun:",pk,autoincrement" json:"id"`
	CustomerName        string      `json:"customer_name"`
	CustomerEmail       string      `json:"customer_email"`
	CustomerPhoneNumber string      `json:"customer_phone_number"`
	InvoiceNumber       string      `json:"invoice_number"`
	WarantyNumber       string      `json:"waranty_number"`
	Detail              string      `json:"detail"`
	InvoiceImage        string      `json:"invoice_image"`
	WarantyCertificate  string      `json:"waranty_certificate"`
	Subject             string      `json:"subject"`
	HelpTopicID         int64       `json:"help_topic_id"`
	HelpTopic           *HelpTopic  `bun:"rel:belongs-to,join:help_topic_id=id" json:"help_topic"`
	DepartmentId        int64       `json:"department_id"`
	Department          *Department `bun:"rel:belongs-to,join:department_id=id" json:"department"`
	StateID             int64       `json:"state_id"`
	State               *State      `bun:"rel:belongs-to,join:state_id=id" json:"state"`
	CreatedByID         int64       `json:"created_by_id"`
	CreatedBy           *User       `bun:"rel:belongs-to,join:created_by_id=id" json:"created_by"`
	CreatedAt           time.Time   `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt           time.Time   `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt           *time.Time  `bun:",nullzero" json:"deleted_at"`
}

func GetAllTickets(db *bun.DB) ([]*Ticket, error) {
	tickets := []*Ticket{}
	ticket := Ticket{}
	err := db.NewSelect().Model(&ticket).
		Relation("HelpTopic").
		Relation("Department").
		Relation("State").
		Relation("CreatedBy").
		Order("id").
		Scan(context.Background(), &tickets)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func GetTicketsByState(db *bun.DB, stateID int64) ([]*Ticket, error) {
	var tickets = []*Ticket{}
	ticket := Ticket{}
	err := db.NewSelect().Model(&ticket).
		Where("state_id = ?", stateID).
		Relation("HelpTopic").
		Relation("Department").
		Relation("State").
		Relation("CreatedBy").
		Scan(context.Background(), &tickets)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func GetTicketByID(db *bun.DB, id int64) (*Ticket, error) {
	ticket := Ticket{}
	err := db.NewSelect().Model(&ticket).Where("ticket.id = ?", id).
		Relation("HelpTopic").
		Relation("Department").
		Relation("State").
		Relation("CreatedBy").
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func CreateTicket(db *bun.DB, ticket *Ticket) (*Ticket, error) {
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()
	ticket.StateID = utils.Open
	_, err := db.NewInsert().Model(ticket).Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return ticket, nil
}
