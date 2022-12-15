package repository

import "github.com/Hanivan/go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
