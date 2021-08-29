package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/kenta4327/example-of-ec-site/app/application/usecase"
	"github.com/kenta4327/example-of-ec-site/app/domain/model"
	"github.com/kenta4327/example-of-ec-site/app/test/mock_repository"
)

func TestBrowseItemUseCase(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockItemRepository := mock_repository.NewMockItemRepository(ctrl)

	itemEntity1 := model.NewItemEntity(1, "商品1", 100)
	mockItemRepository.EXPECT().FindItem(1).Return(*itemEntity1)
	itemEntity2 := model.NewItemEntity(2, "商品2", 200)
	mockItemRepository.EXPECT().FindItem(2).Return(*itemEntity2)

	browseItemUseCase := usecase.NewBrowseItemUseCase(mockItemRepository)

	// ケース1 ： ID1で検索
	browseItemDto1 := browseItemUseCase.BrowseItem(1)
	if browseItemDto1.Id() != 1 || browseItemDto1.Name() != "商品1" || browseItemDto1.Price() != 100 {
		t.Error("ID1での検索が期待と違う")
	}
	// ケース2 ： ID2で検索
	browseItemDto2 := browseItemUseCase.BrowseItem(2)
	if browseItemDto2.Id() != 2 || browseItemDto2.Name() != "商品2" || browseItemDto2.Price() != 200 {
		t.Error("ID2での検索が期待と違う")
	}
}
