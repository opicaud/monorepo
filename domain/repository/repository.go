package repository

import "example2/domain/valueobject"

type Repository interface {
	Save(shape valueobject.Shape) error
}
