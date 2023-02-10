package http

import (
	"github.com/aguspray001/upload-service/usecase"
	"github.com/labstack/echo/v4"
)

type uploadFileHandler struct {
	uploadFileSvc usecase.UploadFileUsecase
}

type UploadFileHandler interface {
	PostMultipleData(ctx echo.Context) error
	PostSingleData(ctx echo.Context) error
	GetFileData(ctx echo.Context) error
}

func NewUploadFileHandler(svc usecase.UploadFileUsecase) UploadFileHandler {
	// constructor handler / usercase
	return &uploadFileHandler{uploadFileSvc: svc}
}
