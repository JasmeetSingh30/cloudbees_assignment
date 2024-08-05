//go:generate mockgen -source=section_service.go -destination=section_service.mockgen.go -package=service

package service

import (
	"fmt"

	"cloudbees_assignmnet.com/ass1/internal/App/custom_errors"
	"cloudbees_assignmnet.com/ass1/internal/App/model"
)

type SectionService interface {
	GetOpenSeats(section *model.Section) int
	GetNextOpenSeat(section model.Section) (*model.Seat, error)
	DeallocateSeat(section model.Section, seatNumber int) error
}

type sectionService struct{}

func NewSectionService() SectionService {
	return &sectionService{}
}

func (a *sectionService) GetOpenSeats(section *model.Section) int {
	openSeats := 0
	for _, seat := range section.Seats {
		if seat.GetPassenger() == nil {
			openSeats += 1
		}
	}

	fmt.Printf("%s has %d available seats\n", section.GetSectionName(), openSeats)

	return openSeats
}

func (a *sectionService) GetNextOpenSeat(section model.Section) (*model.Seat, error) {
	for _, seat := range section.GetSeats() {
		if seat.GetPassenger() == nil {
			return seat, nil
		}
	}
	return nil, customErrors.NotFound("no more seats available in section %s", section.GetSectionName())
}

func (a *sectionService) DeallocateSeat(section model.Section, seatNumber int) error {
	for _, seat := range section.GetSeats() {
		if seat.GetNumber() == seatNumber {
			seat.RemovePassenger()
			return nil
		}
	}
	return customErrors.NotFound("no seat available in section %s with number %d", section.GetSectionName(), seatNumber)
}
