//go:generate mockgen -source=train_service.go -destination=train_service.mockgen.go -package=service

package service

import (
	"fmt"

	customErrors "cloudbees_assignmnet.com/ass1/internal/App/custom_errors"
	"cloudbees_assignmnet.com/ass1/internal/App/local_repo"
	"cloudbees_assignmnet.com/ass1/internal/App/model"
)

type TrainService interface {
	AutoAssignSeat(train *model.Train, passenger *model.User, ticket *model.Ticket) error
	NewTrain(from string, to string, price float32) *model.Train
	InitTrains()
	GetSeatPassengerList(train model.Train, sectionReq string) []model.Seat
	DeallocateSeat(train *model.Train, sectionName string, seatNo int) error
	ModifySeatInSection(train model.Train, passenger *model.User, sectionReq string, seatNoReq int, seatNoExist int) error
}

type trainService struct {
	trainRepo      *local_repo.TrainRepository
	sectionService SectionService
}

func NewTrainService(trainRepo *local_repo.TrainRepository, sectionService SectionService) TrainService {
	return &trainService{trainRepo: trainRepo, sectionService: sectionService}
}

func (s *trainService) NewTrain(from string, to string, price float32) *model.Train {
	return s.trainRepo.Create(from, to, price)
}

func (s *trainService) InitTrains() {
	fmt.Println("Initializing trains: Mumbai -> Kolkata, Bangalore -> New Delhi, New Delhi -> Mumbai")
	s.NewTrain("Mumbai", "Kolkata", 50.0)
	s.NewTrain("Bangalore", "New Delhi", 68.0)
	s.NewTrain("New Delhi", "Mumbai", 40.0)
}

func (s *trainService) AutoAssignSeat(train *model.Train, passenger *model.User, ticket *model.Ticket) error {
	seat, err := s.getNextAvailableSeat(*train, ticket)
	if err != nil {
		return err
	}

	ticket.SeatNo = seat.GetNumber()

	fmt.Printf("assigning passenger to seat %d in %s \n", seat.GetNumber(), seat.GetSection().GetSectionName())
	return seat.SetPassenger(passenger)
}

func (s *trainService) getNextAvailableSeat(train model.Train, ticket *model.Ticket) (*model.Seat, error) {
	section, err := s.getNextAvailableSection(train)
	if err != nil {
		return nil, err
	}

	ticket.Section = section.GetSectionName()

	seats := section.GetSeats()
	for _, seat := range seats {
		if seat.GetPassenger() == nil {
			return seat, nil
		}
	}

	fmt.Printf("no available seats in %s found\n", section.GetSectionName())
	return nil, customErrors.NotFound("the train is fully booked")
}

func (s *trainService) getNextAvailableSection(train model.Train) (*model.Section, error) {
	sections := s.getAvailableSections(train)
	var nextSection model.Section
	mostAvailableSeats := 0

	for _, section := range sections {
		availableSeats := 0
		seats := section.GetSeats()

		fmt.Printf("Current Section %s", section.GetSectionName())

		for _, seat := range seats {
			if seat.GetPassenger() == nil {
				availableSeats += 1
			}
		}
		fmt.Println()

		if availableSeats > mostAvailableSeats {
			nextSection = section
			mostAvailableSeats = availableSeats
		}
		fmt.Printf("Seats available %d, most available %d", availableSeats, mostAvailableSeats)
	}

	if &nextSection == nil {
		fmt.Println("no available sections found")
		return nil, customErrors.NotFound("the train is fully booked")
	}

	fmt.Printf("Section returned %s", nextSection.SectionName)

	return &nextSection, nil
}

func (s *trainService) getAvailableSections(train model.Train) []model.Section {
	allSections := train.GetSections()
	availableSections := []model.Section{}

	for _, section := range allSections {
		if s.sectionService.GetOpenSeats(section) > 0 {
			availableSections = append(availableSections, *section)
		}
	}

	return availableSections
}

func (s *trainService) GetSeatPassengerList(train model.Train, sectionReq string) []model.Seat {
	allSections := train.GetSections()
	seats := make([]model.Seat, 0)
	for _, section := range allSections {
		if sectionReq == section.GetSectionName() {
			for _, seat := range section.Seats {
				seats = append(seats, *seat)
			}
		}
	}
	return seats
}

func (s *trainService) ModifySeatInSection(train model.Train, passenger *model.User, sectionReq string, seatNoReq int, seatNoExist int) error {
	for _, section := range train.GetSections() {
		if sectionReq == section.GetSectionName() {
			for _, seat := range section.Seats {
				if seat.GetNumber() == seatNoReq {
					if seat.Passenger == nil {
						seat.SetPassenger(passenger)
						s.sectionService.DeallocateSeat(*section, seatNoExist)
						return nil
					} else {
						return customErrors.Exists("Passenger already exits at Seat No %d", seatNoReq)
					}
				}
			}
		}
	}
	return customErrors.InvalidInput("Mentioned Seat and Section doesnot Exist, please enter valid inputs")
}

// DeallocateSeat implements TrainService.
func (s *trainService) DeallocateSeat(train *model.Train, sectionName string, seatNo int) error {
	for _, section := range train.Sections {
		if section.GetSectionName() == sectionName {
			error := s.sectionService.DeallocateSeat(*section, seatNo)
			if error != nil {
				return error
			}
		}
	}
	return customErrors.InvalidInput("Invalid details provided, please try with correct details")
}
