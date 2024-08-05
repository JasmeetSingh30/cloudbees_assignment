package server

import (
	"context"

	service "cloudbees_assignmnet.com/ass1/internal/App/services"
	"cloudbees_assignmnet.com/ass1/internal/probuf/probuf_generated"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerServiceMapper struct {
	probuf_generated.UnimplementedRailwayServiceServer
	ticketService *service.TicketService
}

func NewServerServiceMapper(service *service.TicketService) probuf_generated.RailwayServiceServer {
	return &ServerServiceMapper{ticketService: service}
}

func (a *ServerServiceMapper) BuyTicket(ctx context.Context, req *probuf_generated.TicketCreateRequest) (*probuf_generated.Receipt, error){
	firstName := req.GetUser().GetFirstName()
	lastName := req.GetUser().GetLastName()
	email := req.GetUser().GetEmail()
	from := req.GetFrom()
	to := req.GetTo()

	return a.ticketService.PurchaseTicket(ctx, firstName, lastName, email, from, to)
}

func (a *ServerServiceMapper) ViewReceipt(ctx context.Context, req *probuf_generated.TicketRequest) (*probuf_generated.Receipt, error) {
	firstName := req.GetUser().GetFirstName()
	lastName := req.GetUser().GetLastName()
	email := req.GetUser().GetEmail()
	ticketId := req.GetTicketId()

	return a.ticketService.ViewReceipt(ctx, firstName, lastName, email, int(ticketId))
}

func (a *ServerServiceMapper) ViewAllocatedSeats(ctx context.Context, req *probuf_generated.TrainSectionAllocated) (*probuf_generated.Section, error) {
	from := req.GetFrom()
	to := req.GetTo()
	section := req.GetSection().String()
	
	return a.ticketService.ViewAllocatedSeats(ctx, from, to, section)
}

func (a *ServerServiceMapper) CancelTicket(ctx context.Context, req *probuf_generated.TicketRequest) (*emptypb.Empty, error) {
	firstName := req.GetUser().GetFirstName()
	lastName := req.GetUser().GetLastName()
	email := req.GetUser().GetEmail()
	ticketId := req.GetTicketId()

	error := a.ticketService.CancelTicket(ctx, firstName, lastName, email, int(ticketId))
	return nil, error
}

func (a *ServerServiceMapper) ModifySeat(ctx context.Context, req *probuf_generated.SeatModificationRequest) (*probuf_generated.Receipt, error) {
	firstName := req.GetUser().GetFirstName()
	lastName := req.GetUser().GetLastName()
	email := req.GetUser().GetEmail()
	ticketId := req.GetTicketId()
	section := req.GetSection().String()
	seat := req.GetSeat()


	return a.ticketService.ModifySeat(ctx, firstName, lastName, email, int(ticketId), section, int(seat))
}

