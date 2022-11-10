package usecase

import "github.com/12ev09/info-demo-backend/app/domain"

type ItemUsecase interface {
	GetItems(contentType int) ([]domain.Item, error)
	PostItem(domain.Item) error
}
