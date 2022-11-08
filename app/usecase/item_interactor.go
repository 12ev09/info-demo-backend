package usecase

import (
	"github.com/12ev09/info-demo-backend/app/domain"
	"github.com/12ev09/info-demo-backend/app/repository"
)

type itemInteractor struct {
	itemRepository repository.ItemRepository
}

func NewItemInteractor(itemRepository repository.ItemRepository) ItemUsecase {
	return &itemInteractor{
		itemRepository: itemRepository,
	}
}

func (i *itemInteractor) GetItems() ([]domain.Item, error) {
	items, err := i.itemRepository.GetItems()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (i *itemInteractor) PostItem(item domain.Item) error {
	if err := i.itemRepository.PostItem(item); err != nil {
		return err
	}
	return nil
}
