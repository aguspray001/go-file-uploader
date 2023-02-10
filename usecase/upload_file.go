package usecase

import (
	"context"

	"github.com/aguspray001/upload-service/models"
)

func (u *uploadFileUsecase) PostFile(ctx context.Context, payload *models.UploadFile) error {
	err := u.uploadFileRepo.UploadFile(ctx, payload)
	if err != nil {
		return err
	}
	return nil
}

func (u *uploadFileUsecase) PostFiles(ctx context.Context, payloads []*models.UploadFile) error {
	err := u.uploadFileRepo.UploadFiles(ctx, payloads)
	if err != nil {
		return err
	}
	return nil
}

func (u *uploadFileUsecase) GetList(ctx context.Context) ([]*models.UploadFile, error) {
	resp, err := u.uploadFileRepo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
