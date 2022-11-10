package repository

import (
	"fmt"
	"strings"

	"github.com/12ev09/info-demo-backend/app/domain"
	"github.com/jmoiron/sqlx"
)

type ItemRepository interface {
	GetItems(contentType int) ([]domain.Item, error)
	PostItem(domain.Item) error
}

type itemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (i *itemRepository) GetItems(contentType int) ([]domain.Item, error) {
	var items []domain.Item

	args := []string{}

	if contentType != 0 {
		args = append(args, fmt.Sprintf("where content_type=%d", contentType))
	}

	rows, err := i.db.Queryx(fmt.Sprintf("SELECT * from items %s", strings.Join(args, " ")))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := domain.Item{}
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (i *itemRepository) PostItem(item domain.Item) error {
	query := "INSERT INTO items (title, isbn, publisher_name,sales_date,content_type) VALUES (:title, :isbn, :publisher_name, :sales_date, :content_type)"
	if _, err := i.db.NamedExec(query, item); err != nil {
		return err
	}

	return nil
}
