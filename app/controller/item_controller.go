package controller

import (
	"net/http"

	"github.com/12ev09/info-demo-backend/app/domain"
	"github.com/12ev09/info-demo-backend/app/usecase"
	"github.com/labstack/echo/v4"
)

type ItemController interface {
	GetItems(ctx echo.Context) error
	PostItem(ctx echo.Context) error
}

type itemController struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemController(itemUsecase usecase.ItemUsecase) ItemController {
	return &itemController{
		itemUsecase: itemUsecase,
	}
}

func (c *itemController) GetItems(ctx echo.Context) error {
	items := []domain.Item{
		{
			Title: "eee",
		},
	}
	// items, err := c.itemUsecase.GetItems()
	// if err != nil {
	// 	return nil
	// }

	return ctx.JSON(http.StatusOK, items)
}

func (c *itemController) PostItem(ctx echo.Context) error {
	return nil
}
