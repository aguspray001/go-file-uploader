package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/aguspray001/upload-service/config/postgres"
	delivery "github.com/aguspray001/upload-service/delivery/http"
	repository "github.com/aguspray001/upload-service/repository/postgres"
	"github.com/aguspray001/upload-service/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := config.PostgresConfig{}
	db, err := conf.ConnectToDB()
	if err != nil {
		panic("Something wrong when connecting the DB")
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "assets")

	uploadFileRepository := repository.NewPostgresUploadFileRepository(db)
	uploadFileUsecase := usecase.NewUploadFileUsecase(uploadFileRepository)
	uploadFileDelivery := delivery.NewUploadFileHandler(uploadFileUsecase)

	routingPrefix := "/api/v1"

	e.POST(fmt.Sprintf(routingPrefix+"/file/single"), uploadFileDelivery.PostSingleData)
	e.POST(fmt.Sprintf(routingPrefix+"/file/multiple"), uploadFileDelivery.PostMultipleData)
	e.GET(fmt.Sprintf(routingPrefix+"/files"), uploadFileDelivery.GetFileData)

	e.GET(fmt.Sprintf(routingPrefix+"/health-checker"), func(c echo.Context) error {
		return c.String(http.StatusOK, "Wellcome to Upload File API!")
	})

	if err := e.Start(":3001"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
