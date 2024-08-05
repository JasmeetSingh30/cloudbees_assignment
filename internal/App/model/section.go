package model

import (
	"gorm.io/gorm"
)

type Section struct {
	gorm.Model

	SectionName string
	Seats map[int]*Seat
}

func NewSection(sectionName string) *Section {
	section := &Section{SectionName: sectionName, Seats: make(map[int]*Seat)}

	for i := 1; i <= 10; i++ {
		section.Seats[i] = NewSeat(i, section)
	}

	return section
}

func (a *Section) GetSectionName() string {
	return a.SectionName
}

func (a *Section) GetSeats() map[int]*Seat {
	return a.Seats
}