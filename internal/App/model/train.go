package model

import "gorm.io/gorm"

type Train struct {
	gorm.Model

	From     string
	To       string
	Price    float32
	Sections []*Section
}

func NewTrain(from string, to string, price float32) *Train {
	train := &Train{From: from, To: to, Price: price}
	sections := []*Section{
		NewSection("SECTION_A"),
		NewSection("SECTION_B"),
	}
	train.Sections = sections

	return train
}

func (a *Train) GetSections() []*Section {
	return a.Sections
}
