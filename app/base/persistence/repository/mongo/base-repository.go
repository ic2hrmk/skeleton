package mongo

import (
	"github.com/ic2hrmk/skeleton/app/base/persistence/model"
	"github.com/ic2hrmk/skeleton/app/base/persistence/repository"
)

type baseRepository struct {
	db string // Any DB connection
}

func NewBaseRepository() repository.BaseRepository {
	return &baseRepository{}
}

func (rcv *baseRepository) FindAll() ([]*model.Base, error) {
	return nil, nil
}


