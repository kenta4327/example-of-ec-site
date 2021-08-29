package model

type ItemEntity struct {
	id    int
	name  string
	price int
}

func NewItemEntity(id int, name string, price int) *ItemEntity {
	return &ItemEntity{
		id:id,
		name:  name,
		price: price,
	}
}

func (itemEntity ItemEntity) Id() int {
	return itemEntity.id
}

func (itemEntity ItemEntity) Name() string {
	return itemEntity.name
}

func (itemEntity ItemEntity) Price() int {
	return itemEntity.price
}
