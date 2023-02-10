package http

import (
	// "context"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aguspray001/upload-service/models"
	"github.com/labstack/echo/v4"
)

func (h *uploadFileHandler) PostMultipleData(ctx echo.Context) error {
	userID := ctx.FormValue("user_id")
	respJson := []*models.UploadFile{}
	form, err := ctx.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["files"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// destination
		timestamp := time.Now().Unix()
		filename := fmt.Sprintf("%d_%s", timestamp, file.Filename)
		dst, err := os.Create("./assets/" + filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		respJson = append(respJson, &models.UploadFile{
			UserID:   userID,
			FileName: file.Filename,
		})
	}
	err = h.uploadFileSvc.PostFiles(context.Background(), respJson)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.ResponseObject{
			Status:  http.StatusBadRequest,
			Message: "Data are unsuccessfully saved in database",
			Data:    &respJson,
		})
	}
	return ctx.JSON(http.StatusOK, &models.ResponseObject{
		Status:  http.StatusOK,
		Message: "Data are successfully saved in database",
		Data:    &respJson,
	})
}

func (h *uploadFileHandler) PostSingleData(ctx echo.Context) error {
	userID := ctx.FormValue("user_id")
	respJson := models.UploadFile{}
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return nil
	}
	defer src.Close()

	// destination
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d_%s", timestamp, file.Filename)

	dst, err := os.Create("./assets/" + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	respJson.UserID = userID
	respJson.FileName = filename

	err = h.uploadFileSvc.PostFile(context.Background(), &respJson)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.ResponseObject{
			Status:  http.StatusBadRequest,
			Message: "Data is unsuccessfully saved in database",
			Data: &models.UploadFile{
				UserID:   userID,
				FileName: file.Filename,
			},
		})
	}
	return ctx.JSON(http.StatusOK, &models.ResponseObject{
		Status:  http.StatusOK,
		Message: "Data is successfully saved in database",
		Data: &models.UploadFile{
			UserID:   userID,
			FileName: file.Filename,
		},
	})
}

func (h *uploadFileHandler) GetFileData(ctx echo.Context) error {
	// get file from db
	data, err := h.uploadFileSvc.GetList(context.Background())

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.ResponseObject{
			Status:  http.StatusBadRequest,
			Message: "Data is unsuccessfully get from database",
			Data:    data,
		})
	}

	return ctx.JSON(http.StatusOK, &models.ResponseObject{
		Status:  http.StatusOK,
		Message: "Data is successfully get from database",
		Data:    data,
	})
}
