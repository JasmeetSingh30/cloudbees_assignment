package local_repo

import (
	"fmt"

	"cloudbees_assignmnet.com/ass1/internal/App/custom_errors"
	"cloudbees_assignmnet.com/ass1/internal/App/model"
)

type PassengerRepository struct {
	passCounter uint
	passengers map[uint]*model.User
}

func NewPassengerRepository() *PassengerRepository {
	return &PassengerRepository{passCounter: 0, passengers: map[uint]*model.User{}}
}

func (r *PassengerRepository) GetByID(id uint) (*model.User, error) {
	passenger, ok := r.passengers[id]
	if !ok {
		return nil, customErrors.NotFound("user %d", id)
	}

	return passenger, nil
}

func (r *PassengerRepository) GetByData(firstName string, lastName string, email string) (*model.User, error) {
	for _, user := range r.passengers {
		if user.FirstName == firstName && user.LastName == lastName && user.Email == email {
			return user, nil
		}
	}

	return nil, customErrors.NotFound("user %s %s", firstName, lastName)
}

func (r *PassengerRepository) FindOrCreate(firstName string, lastName string, email string) *model.User {
	if user, _ := r.GetByData(firstName, lastName, email); user != nil {
		return user
	}

	passenger := model.NewUser(firstName, lastName, email)
	// auto-increment the ID
	passenger.ID = r.generateID()

	r.passengers[passenger.ID] = passenger
	r.passCounter = passenger.ID

	return passenger
}

func (r *PassengerRepository) generateID() uint {
	return r.passCounter + 1
}

func (r *PassengerRepository) RemoveTicket(firstName string, lastName string, email string, ticketId int) error {
	passenger, error := r.GetByData(firstName, lastName, email)
	if error != nil{
		return error
	}

	for index, ticket := range passenger.Tickets {
		fmt.Println("Index", index)
		if ticket.ID == uint(ticketId){
			passenger.Tickets = append(passenger.Tickets[:index], passenger.Tickets[index+1:]...)
			return nil
		}
	}
	return customErrors.NotFound("Ticket with ID %d not found", ticketId)
}

func (r *PassengerRepository) UpdateTicket(ticketModified model.Ticket) error {
	passenger, err := r.GetByID(ticketModified.User.ID)
	if err != nil {
		return err
	}

	for _, ticket := range passenger.Tickets{
		if ticket.GetID() == ticketModified.GetID(){
			ticket.Section = ticketModified.Section
			ticket.SeatNo = ticketModified.SeatNo
			return nil
		}
	}
	return customErrors.NotFound("Ticket with ID %d does not exist", ticketModified.GetID())
}
