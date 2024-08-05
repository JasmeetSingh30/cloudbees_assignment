package service

import (
	"context"

	customErrors "cloudbees_assignmnet.com/ass1/internal/App/custom_errors"
	LocalRepo "cloudbees_assignmnet.com/ass1/internal/App/local_repo"
	"cloudbees_assignmnet.com/ass1/internal/probuf/probuf_generated"
)

type TicketService struct {
	ticketRepo    *LocalRepo.TicketRepository
	trainRepo     *LocalRepo.TrainRepository
	trainService  TrainService
	passengerRepo *LocalRepo.PassengerRepository
}

func NewTicketService(ticketRepo *LocalRepo.TicketRepository, passengerRepo *LocalRepo.PassengerRepository, trainRepo *LocalRepo.TrainRepository, trainService TrainService) *TicketService {
	return &TicketService{ticketRepo: ticketRepo, passengerRepo: passengerRepo, trainRepo: trainRepo, trainService: trainService}
}

func (s *TicketService) PurchaseTicket(ctx context.Context, firstName string, lastName string, email string, from string, to string) (*probuf_generated.Receipt, error) {
	train, err := s.trainRepo.Find(from, to)
	if err != nil {
		return nil, err
	}

	user := s.passengerRepo.FindOrCreate(firstName, lastName, email)

	ticket, err := s.ticketRepo.Create(ctx, user, train)
	if err != nil {
		return nil, err
	}

	err = s.trainService.AutoAssignSeat(train, user, ticket)
	if err != nil {
		s.ticketRepo.Delete(ctx, ticket.GetID())
		return nil, err
	}

	return ticket.GenerateReceipt(), nil
}

func (s *TicketService) ViewReceipt(ctx context.Context, firstName string, lastName string, email string, ticketId int) (*probuf_generated.Receipt, error) {

	user, error := s.passengerRepo.GetByData(firstName, lastName, email)
	if error != nil {
		return nil, error
	}

	for _, ticket := range user.Tickets {
		if ticket.ID == uint(ticketId) {
			return ticket.GenerateReceipt(), nil
		}
	}

	return nil, customErrors.NotFound("Ticket Does not exist with Id: %d", ticketId)
}

func (s *TicketService) ViewAllocatedSeats(ctx context.Context, from string, to string, sectionReq string) (*probuf_generated.Section, error) {
	train, err := s.trainRepo.Find(from, to)
	if err != nil {
		return nil, err
	}

	seats := s.trainService.GetSeatPassengerList(*train, sectionReq)

	resultSeats := make([]*probuf_generated.Seat, 0)
	for _, seat := range seats {
		if seat.Passenger != nil {
			resultSeats = append(resultSeats, &probuf_generated.Seat{
				User: &probuf_generated.User{
					FirstName: seat.Passenger.FirstName,
					LastName:  seat.Passenger.LastName,
					Email:     seat.Passenger.Email,
				},
				SeatNo: int32(seat.GetNumber()),
			})
		} else {
			resultSeats = append(resultSeats, &probuf_generated.Seat{
				User:   nil,
				SeatNo: int32(seat.GetNumber()),
			})

		}
	}
	return &probuf_generated.Section{
		SectionName: sectionReq,
		Seats:       resultSeats,
	}, nil
}

func (s *TicketService) CancelTicket(ctx context.Context, firstName string, lastName string, email string, ticketId int) error {
	ticket := s.ticketRepo.GetTicketByID(uint(ticketId))

	error := s.passengerRepo.RemoveTicket(firstName, lastName, email, ticketId)
	if error != nil {
		return error
	}
	s.trainService.DeallocateSeat(ticket.GetTrain(), ticket.Section, ticket.SeatNo)
	s.ticketRepo.Delete(ctx, uint(ticketId))
	return nil
}

func (s *TicketService) ModifySeat(ctx context.Context, firstName string, lastName string, email string, ticketId int, reqSec string, reqSeatNo int) (*probuf_generated.Receipt, error) {
	passenger, err := s.passengerRepo.GetByData(firstName, lastName, email)
	if err != nil {
		return nil, err
	}

	ticket := s.ticketRepo.GetTicketByID(uint(ticketId))
	if ticket == nil {
		return nil, customErrors.NotFound("Ticket with Id %d does not exist", ticketId)
	}

	err = s.trainService.ModifySeatInSection(*ticket.Train, passenger, reqSec, reqSeatNo, ticket.SeatNo)
	if err != nil {
		return nil, err
	}

	ticket.SeatNo = reqSeatNo
	ticket.Section = reqSec

	s.passengerRepo.UpdateTicket(*ticket)
	return ticket.GenerateReceipt(), nil
}
