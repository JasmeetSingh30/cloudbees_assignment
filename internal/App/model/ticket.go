package model

import (
	"cloudbees_assignmnet.com/ass1/internal/probuf/probuf_generated"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model

	User    *User
	Train   *Train
	Price   float32
	Section string
	SeatNo  int
}

func NewTicket(user *User, train *Train, price float32) *Ticket {
	return &Ticket{User: user, Train: train, Price: price}
}

func (a *Ticket) GetID() uint {
	return a.ID
}

func (a *Ticket) GetSeatNo() int {
	return a.SeatNo
}

func (a *Ticket) GetTrain() *Train {
	return a.Train
}

func (a *Ticket) GenerateReceipt() *probuf_generated.Receipt {
	return &probuf_generated.Receipt{
		User: &probuf_generated.User{
			FirstName: a.User.FirstName,
			LastName:  a.User.LastName,
			Email:     a.User.Email,
		},
		From:  a.Train.From,
		To:    a.Train.To,
		Price: a.Train.Price,
		TicketId: int32(a.GetID()),
		SeatNo: int32(a.GetSeatNo()),
		Section: a.Section,
	}
}
