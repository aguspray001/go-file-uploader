package repository

import (
	"context"
	"fmt"

	"github.com/aguspray001/upload-service/models"
)

func (r *uploadFileRepository) UploadFile(ctx context.Context, payload *models.UploadFile) error {
	if payload == nil {
		fmt.Println("there is no file for uploading")
	}
	result := r.db.Create(&payload)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *uploadFileRepository) UploadFiles(ctx context.Context, payload []*models.UploadFile) error {
	if payload == nil {
		fmt.Println("there is no file for uploading")
	}
	result := r.db.Create(&payload)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *uploadFileRepository) GetList(ctx context.Context) ([]*models.UploadFile, error) {
	var files []*models.UploadFile
	resp := r.db.Find(&files)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return files, nil
}
