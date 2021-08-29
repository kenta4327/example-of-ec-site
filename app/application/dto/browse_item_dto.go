package dto

import (
	"github.com/kenta4327/example-of-ec-site/app/domain/model"
)

type BrowseItemDto struct {
	id    int
	name  string
	price int
}

func NewBrowseItemDto(itemEntity model.ItemEntity) *BrowseItemDto {
	return &BrowseItemDto{
		id:    itemEntity.Id(),
		name:  itemEntity.Name(),
		price: itemEntity.Price(),
	}
}

func (browseItemDto BrowseItemDto) Id() int {
	return browseItemDto.id
}

func (browseItemDto BrowseItemDto) Name() string {
	return browseItemDto.name
}

func (browseItemDto BrowseItemDto) Price() int {
	return browseItemDto.price
}
