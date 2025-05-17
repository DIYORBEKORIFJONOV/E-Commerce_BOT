package minao1

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/config"
	"github.com/minio/minio-go/v7"
)

type (
	FileStorage struct {
		cfg          *config.Config
		minio_client *minio.Client
	}
)

func NewFileStorage(cfg *config.Config, minio_client *minio.Client) *FileStorage {
    fs := &FileStorage{
        cfg:          cfg,
        minio_client: minio_client,
    }

    bucketName := "products"
    ctx := context.Background()
    exists, err := minio_client.BucketExists(ctx, bucketName)
    if err != nil {
        log.Fatalf("Failed to check if bucket exists: %v", err)
    }
    if !exists {
        err = minio_client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
        if err != nil {
            log.Fatalf("Failed to create bucket: %v", err)
        }
        log.Printf("Bucket %s created successfully", bucketName)
    }

    return fs
}

func (s *FileStorage) UploadFile(file *multipart.FileHeader) (string, error) {

	// Open file
	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("cannot open file: %s", err.Error())
	}
	defer f.Close()

	// Generate unique filename
	filename := fmt.Sprintf("%d", time.Now().UnixMilli())

	// Upload file to MinIO
	_, err = s.minio_client.PutObject(
		context.Background(),
		"products",
		filename,
		f,
		file.Size,
		minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %s", err.Error())
	}

	return filename, nil
}

func (s *FileStorage) UploadFileFromBytes(data []byte, contentType string) (string, error) {
	// Create a reader from the raw bytes
	reader := bytes.NewReader(data)
	filename := fmt.Sprintf("%d", time.Now().UnixMilli())

	// Upload file to MinIO
	_, err := s.minio_client.PutObject(
		context.Background(),
		"products",
		filename,
		reader,
		int64(len(data)),
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload file from bytes: %s", err.Error())
	}

	return s.GetFile(filename)
}
func (s *FileStorage) GetFile(filename string) (string, error) {
    _, err := s.minio_client.StatObject(context.Background(), "products", filename, minio.StatObjectOptions{})
    if err != nil {
        return "", fmt.Errorf("file not found: %s", err.Error())
    }

    expiry := time.Hour * 24
    url, err := s.minio_client.PresignedGetObject(context.Background(), "products", filename, expiry, nil)
    if err != nil {
        return "", fmt.Errorf("failed to get file: %s", err.Error())
    }

    finalURL := url.String()
    log.Printf("Сгенерированный URL: %s", finalURL)
    return finalURL, nil
}
func (s *FileStorage) DeleteFile(filename string) error {
	// Check if the file exists before attempting deletion
	_, err := s.minio_client.StatObject(context.Background(),"products", filename, minio.StatObjectOptions{})
	if err != nil {
		return fmt.Errorf("file not found: %s", err.Error())
	}

	// Delete file from MinIO
	err = s.minio_client.RemoveObject(context.Background(), "products", filename, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file: %s", err.Error())
	}

	return nil
}
