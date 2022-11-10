package usecase

import "github.com/12ev09/info-demo-backend/app/domain"

type ItemUsecase interface {
	GetItems(contentType int) ([]OutputItem, error)
	PostItem(domain.Item) error
	DeleteItem(id int) error
	UpdateItem(input UpdateItemInput) error
}

type OutputItem struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Isbn          string `json:"isbn"`
	PublisherName string `json:"publisher_name"`
	SalesDate     string `json:"sales_date"`
	ContentType   string `json:"content_type"`
}

type UpdateItemInput struct {
	ID            uint
	Title         string
	Isbn          string
	PublisherName string
	SalesDate     string
	ContentType   int
}
