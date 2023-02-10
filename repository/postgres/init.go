package repository

import (
	"context"

	"github.com/aguspray001/upload-service/models"
	"gorm.io/gorm"
)

type uploadFileRepository struct {
	db *gorm.DB
}

type UploadFileRepository interface {
	UploadFile(ctx context.Context, payload *models.UploadFile) error
	UploadFiles(ctx context.Context, payload []*models.UploadFile) error
	GetList(ctx context.Context) ([]*models.UploadFile, error)
}

func NewPostgresUploadFileRepository(db *gorm.DB) UploadFileRepository {
	return &uploadFileRepository{db: db}
}
