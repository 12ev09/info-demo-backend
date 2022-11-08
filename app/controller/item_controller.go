package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

	fmt.Print("hi")
	// items, err := c.itemUsecase.GetItems()
	// if err != nil {
	// 	return nil
	// }

	return ctx.JSON(http.StatusOK, items)
}

func (c *itemController) PostItem(ctx echo.Context) error {
	item := domain.Item{
		CreatedAt:     time.Now(),
		Title:         "title",
		Isbn:          "isbn",
		PublisherName: "p",
		SalesDate:     "s",
		ContentType:   1,
	}

	log.Print("hi")

	if err := c.itemUsecase.PostItem(item); err != nil {
		return err
	}
	return ctx.String(http.StatusOK, "post accepted")
}
