package repository

import (
	"github.com/12ev09/info-demo-backend/app/domain"
	"github.com/jmoiron/sqlx"
)

type ItemRepository interface {
	GetItems() ([]domain.Item, error)
	PostItem(domain.Item) error
}

type itemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (i *itemRepository) GetItems() ([]domain.Item, error) {
	return nil, nil
}

func (i *itemRepository) PostItem(domain.Item) error {
	return nil
}
