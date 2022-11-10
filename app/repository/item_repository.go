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
	DeleteItem(id int) error
	UpdateItem(domain.Item) error
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
		args = append(args, fmt.Sprintf("content_type=%d", contentType))
	}

	if len(args) > 0 {
		rows, err := i.db.Queryx(fmt.Sprintf("SELECT * from items where %s", strings.Join(args, " and ")))
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

	rows, err := i.db.Queryx("SELECT * from items")
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
	query := "INSERT INTO items (author, title, isbn, publisher_name,sales_date,content_type) VALUES (:author, :title, :isbn, :publisher_name, :sales_date, :content_type)"
	if _, err := i.db.NamedExec(query, item); err != nil {
		return err
	}

	return nil
}

func (i *itemRepository) UpdateItem(item domain.Item) error {
	query := "UPDATE items SET author = :author, title = :title, isbn = :isbn, publisher_name = :publisher_name,sales_date = :sales_date, content_type = :content_type where id = :id"
	if _, err := i.db.NamedExec(query, item); err != nil {
		return err
	}

	return nil
}

func (i *itemRepository) DeleteItem(id int) error {
	query := "DELETE from items where id = :id"
	if _, err := i.db.NamedExec(query, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
