package local_repo

import (
	customErrors "cloudbees_assignmnet.com/ass1/internal/App/custom_errors"
	"cloudbees_assignmnet.com/ass1/internal/App/model"
)

type TrainRepository struct {
	maxID  uint
	trains map[uint]*model.Train
}

func NewTrainRepository() *TrainRepository {
	trains := make(map[uint]*model.Train)

	return &TrainRepository{maxID: 0, trains: trains}
}

func (r *TrainRepository) GetByID(id uint) (*model.Train, error) {
	train, ok := r.trains[id]
	if !ok {
		return nil, customErrors.NotFound("train with id %d", id)
	}

	return train, nil
}

func (r *TrainRepository) Find(from string, to string) (*model.Train, error) {
	for _, train := range r.trains {
		if train.From == from && train.To == to {
			return train, nil
		}
	}

	return nil, customErrors.NotFound("train from %s to %s not found", from, to)
}

func (r *TrainRepository) Create(from string, to string, price float32) *model.Train {
	if existingTrain, _ := r.Find(from, to); existingTrain != nil {
		return existingTrain
	}

	train := model.NewTrain(from, to, price)
	train.ID = r.generateID()

	r.trains[train.ID] = train
	r.maxID = train.ID

	return train
}

func (r *TrainRepository) generateID() uint {
	return r.maxID + 1
}
