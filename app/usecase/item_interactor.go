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

func (i *itemInteractor) GetItems(contentType int) ([]OutputItem, error) {
	items, err := i.itemRepository.GetItems(contentType)
	if err != nil {
		return nil, err
	}

	output := []OutputItem{}

	for _, item := range items {
		output = append(output, OutputItem{
			ID:            item.ID,
			Author:        item.Author,
			Title:         item.Title,
			Isbn:          item.Isbn,
			PublisherName: item.PublisherName,
			SalesDate:     item.SalesDate,
			ContentType:   domain.ContentTypeString(item.ContentType),
		})
	}

	return output, nil
}

func (i *itemInteractor) PostItem(item domain.Item) error {
	if err := i.itemRepository.PostItem(item); err != nil {
		return err
	}
	return nil
}

func (i *itemInteractor) DeleteItem(id int) error {
	if err := i.itemRepository.DeleteItem(id); err != nil {
		return err
	}
	return nil
}

func (i *itemInteractor) UpdateItem(input UpdateItemInput) error {
	item := domain.Item{
		ID:            input.ID,
		Title:         input.Title,
		Isbn:          input.Isbn,
		PublisherName: input.PublisherName,
		SalesDate:     input.SalesDate,
		ContentType:   input.ContentType,
	}

	if err := i.itemRepository.UpdateItem(item); err != nil {
		return err
	}
	return nil
}
