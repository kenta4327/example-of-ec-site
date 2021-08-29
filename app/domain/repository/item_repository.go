package repository

import (
	"github.com/kenta4327/example-of-ec-site/app/domain/model"
)

type ItemRepository interface {
	FindItem(id int) model.ItemEntity
}
