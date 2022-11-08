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

func (i *itemRepository) PostItem(item domain.Item) error {
	query := "INSERT INTO items (created_at, title, isbn, publisher_name,sales_date,content_type) VALUES (:created_at, :title, :isbn, :publisher_name, :sales_date, :content_type)"
	if _, err := i.db.NamedExec(query, item); err != nil {
		return err
	}

	return nil
}
