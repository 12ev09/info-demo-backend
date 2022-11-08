package router

import (
	"github.com/12ev09/info-demo-backend/app/controller"
	"github.com/12ev09/info-demo-backend/app/db"
	"github.com/12ev09/info-demo-backend/app/repository"
	"github.com/12ev09/info-demo-backend/app/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	itemRepostiory := repository.NewItemRepository(db.DB)

	itemUsecase := usecase.NewItemInteractor(itemRepostiory)

	itemController := controller.NewItemController(itemUsecase)

	// ルートを設定
	e.GET("/items", itemController.GetItems)
	//e.POST("/items")

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(e.Start(":8080"))
	}
}
