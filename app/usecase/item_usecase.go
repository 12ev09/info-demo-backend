package usecase

import "github.com/12ev09/info-demo-backend/app/domain"

type ItemUsecase interface {
	GetItems(contentType int) ([]domain.Item, error)
	PostItem(domain.Item) error
	DeleteItem(id int) error
	UpdateItem(input UpdateItemInput) error
}

type UpdateItemInput struct {
	ID            uint
	Title         string
	Isbn          string
	PublisherName string
	SalesDate     string
	ContentType   string
}
