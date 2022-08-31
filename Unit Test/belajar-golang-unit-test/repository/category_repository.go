package repository

import (
	"belajar-golang-unit-test/entity"
)

type CategoryRepository interface {
	FindId(id string) *entity.Category
}
