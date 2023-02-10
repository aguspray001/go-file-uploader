package usecase

import (
	"context"

	"github.com/aguspray001/upload-service/models"
	repository "github.com/aguspray001/upload-service/repository/postgres"
)

type uploadFileUsecase struct {
	uploadFileRepo repository.UploadFileRepository
}

type UploadFileUsecase interface {
	PostFile(ctx context.Context, payload *models.UploadFile) error
	PostFiles(ctx context.Context, payload []*models.UploadFile) error
	GetList(ctx context.Context) ([]*models.UploadFile, error)
}

func NewUploadFileUsecase(repo repository.UploadFileRepository) UploadFileUsecase {
	return &uploadFileUsecase{
		uploadFileRepo: repo,
	}
}
