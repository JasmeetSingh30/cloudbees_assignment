package model

import (
	customErrors "cloudbees_assignmnet.com/ass1/internal/App/custom_errors"
	"gorm.io/gorm"
)

type Seat struct {
	gorm.Model

	Section   *Section
	Number    int
	Passenger *User
}

func NewSeat(number int, section *Section) *Seat {
	return &Seat{Number: number, Section: section}
}

func (s *Seat) SetPassenger(user *User) error {
	if s.Passenger != nil {
		return customErrors.Unavailable("seat %d in %s is alredy assigned", s.Number, s.Section.SectionName)
	}
	s.Passenger = user
	return nil
}

func (a *Seat) GetNumber() int {
	return a.Number
}

func (a *Seat) GetPassenger() *User {
	return a.Passenger
}

func (a *Seat) GetSection() *Section {
	return a.Section
}

func (a *Seat) RemovePassenger() {
	a.Passenger = nil
}
