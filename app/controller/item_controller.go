package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/12ev09/info-demo-backend/app/domain"
	"github.com/12ev09/info-demo-backend/app/usecase"
	"github.com/labstack/echo/v4"
)

type ItemController interface {
	GetItems(ctx echo.Context) error
	PostItem(ctx echo.Context) error
	DeleteItem(ctx echo.Context) error
	UpdateItem(ctx echo.Context) error
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
	type params struct {
		ContentType int `query:"type"`
	}

	p := params{}
	if err := ctx.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Print(p.ContentType)

	items, err := c.itemUsecase.GetItems(p.ContentType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, items)
}

func (c *itemController) PostItem(ctx echo.Context) error {
	type form struct {
		Author        string `json:"author"`
		Title         string `json:"title"`
		Isbn          string `json:"isbn"`
		PublisherName string `json:"publisher_name"`
		SalesDate     string `json:"sales_date"`
		ContentType   int    `json:"type"`
	}

	f := form{}
	if err := ctx.Bind(&f); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item := domain.Item{
		Author:        f.Author,
		Title:         f.Title,
		Isbn:          f.Isbn,
		PublisherName: f.PublisherName,
		SalesDate:     f.SalesDate,
		ContentType:   f.ContentType,
	}

	if err := c.itemUsecase.PostItem(item); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, "post accepted")
}

func (c *itemController) DeleteItem(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.itemUsecase.DeleteItem(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, fmt.Sprintf("record %d is deleted", id))
}

func (c *itemController) UpdateItem(ctx echo.Context) error {
	type form struct {
		Author        string `json:"author"`
		Title         string `json:"title"`
		Isbn          string `json:"isbn"`
		PublisherName string `json:"publisher_name"`
		SalesDate     string `json:"sales_date"`
		ContentType   int    `json:"type"`
	}

	f := form{}
	if err := ctx.Bind(&f); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updateItemInput := usecase.UpdateItemInput{
		Title:         f.Title,
		Isbn:          f.Isbn,
		PublisherName: f.PublisherName,
		SalesDate:     f.SalesDate,
		ContentType:   f.ContentType,
	}

	if err := c.itemUsecase.UpdateItem(updateItemInput); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, fmt.Sprintf("record %d is updated", id))
}
