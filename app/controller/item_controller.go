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
	items, err := c.itemUsecase.GetItems()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, items)
}

func (c *itemController) PostItem(ctx echo.Context) error {
	type form struct {
		title         string `json:"title"`
		isbn          string `json:"isbn"`
		publisherName string `json:"publisher_name"`
		salesDate     string `json:"sales_date"`
		contentType   int    `json:"content_type"`
	}

	f := form{}
	if err := ctx.Bind(&f); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item := domain.Item{
		//CreatedAt:     time.Now(),
		Title:         f.title,
		Isbn:          f.isbn,
		PublisherName: f.publisherName,
		SalesDate:     f.salesDate,
		ContentType:   f.contentType,
	}

	if err := c.itemUsecase.PostItem(item); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, "post accepted")
}
