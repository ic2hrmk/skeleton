package repository

import (
	"github.com/ic2hrmk/skeleton/app/base/persistence/model"
	"github.com/pkg/errors"
)

var (
	ErrBaseNotFound = errors.New("errBaseNotFound")
)

type BaseRepository interface {
	FindAll() ([]*model.Base, error)
}
