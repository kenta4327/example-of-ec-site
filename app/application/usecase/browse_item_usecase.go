package usecase

import (
	"github.com/kenta4327/example-of-ec-site/app/domain/repository"
	"github.com/kenta4327/example-of-ec-site/app/application/dto"
)

type BrowseItemUseCase struct {
	itemRepository repository.ItemRepository
}

func NewBrowseItemUseCase(itemRepository repository.ItemRepository) *BrowseItemUseCase {
	return &BrowseItemUseCase {
		itemRepository: itemRepository,
	}
}

func (browseItemUseCase BrowseItemUseCase) BrowseItem(itemId int) *dto.BrowseItemDto {
    itemEntity := browseItemUseCase.itemRepository.FindItem(itemId)
	return dto.NewBrowseItemDto(itemEntity)
}
