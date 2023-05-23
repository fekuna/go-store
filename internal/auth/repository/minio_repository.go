package repository

import (
	"context"
	"fmt"

	"github.com/fekuna/go-store/internal/auth"
	"github.com/fekuna/go-store/internal/models"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

// Auth Minio S3 repository
type authMinioRepository struct {
	client *minio.Client
}

// Auth Minio S3 Repository constructor
func NewAuthMinioRepository(minioClient *minio.Client) auth.MinioRepository {
	return &authMinioRepository{client: minioClient}
}

// Upload file to Minio
func (r *authMinioRepository) PutObject(ctx context.Context, input models.UploadInput) (*minio.UploadInfo, error) {
	// TODO: Tracing
	options := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	uploadInfo, err := r.client.PutObject(ctx, input.BucketName, r.generateFileName(input.Name), input.File, input.Size, options)
	if err != nil {
		fmt.Println("Mashok", err)
		return nil, errors.Wrap(err, "authAWSRepository.FileUpload.PutObject")
	}
	return &uploadInfo, err
}

// Download file from minio
func (r *authMinioRepository) GetObject(ctx context.Context, bucket string, fileName string) (*minio.Object, error) {
	// TODO: Tracing

	object, err := r.client.GetObject(ctx, bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "authAWSRepository.FileDownload.GetObject")
	}

	return object, nil
}

// Delete file from Minio
func (r *authMinioRepository) RemoveObject(ctx context.Context, bucket string, fileName string) error {
	// TODO: Tracing

	if err := r.client.RemoveObject(ctx, bucket, fileName, minio.RemoveObjectOptions{}); err != nil {
		return errors.Wrap(err, "authAWSRepository.RemoveObject")
	}

	return nil
}

func (r *authMinioRepository) generateFileName(fileName string) string {
	uid := uuid.New().String()
	return fmt.Sprintf("%s-%s", uid, fileName)
}
