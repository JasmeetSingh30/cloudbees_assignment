package local_repo

import (
	"context"
	"fmt"

	customErrors "cloudbees_assignmnet.com/ass1/internal/App/custom_errors"
	"cloudbees_assignmnet.com/ass1/internal/App/model"
)

type TicketRepository struct {
	ticketCounter uint
	tickets       map[uint]*model.Ticket
	userRepo      *PassengerRepository
}

func NewTicketRepository(userRepo *PassengerRepository) *TicketRepository {
	return &TicketRepository{tickets: map[uint]*model.Ticket{}, userRepo: userRepo}
}

func (r *TicketRepository) Create(ctx context.Context, user *model.User, train *model.Train) (*model.Ticket, error) {
	for _, ticket := range user.Tickets {
		if ticket.Train.From == train.From && ticket.Train.To == train.To {
			return nil, customErrors.Exists("Similar ticket already exist with ticket Id: %d", ticket.ID)
		}
	}

	ticket := model.NewTicket(user, train, train.Price)
	ticket.ID = r.generateID()

	r.tickets[ticket.ID] = ticket
	r.ticketCounter = ticket.ID

	fmt.Printf("Booked ticket from %s to %s  for %s %s at %f with id %d\n", train.From, train.To, user.FirstName, user.LastName, train.Price, ticket.ID)

	user.Tickets = append(user.Tickets, ticket)

	return ticket, nil
}

func (r *TicketRepository) Delete(ctx context.Context, id uint) {
	fmt.Printf("deleting ticket %d\n", id)
	delete(r.tickets, id)


}

func (r *TicketRepository) generateID() uint {
	return r.ticketCounter + 1
}

func (r *TicketRepository) GetTicketByID(ticketId uint) *model.Ticket{
	val, ok := r.tickets[ticketId]
	if ok {
		return val
	}
	return nil
}

func (r *TicketRepository) UpdateTicket(ticketModified model.Ticket) error {
	ticket := r.GetTicketByID(ticketModified.GetID())
	if &ticket == nil{
		return customErrors.NotFound("Ticket with ID %d does not exist", ticketModified.GetID())
	}
	ticket.SeatNo = ticketModified.SeatNo
	ticket.Section = ticketModified.Section
	return nil
}
