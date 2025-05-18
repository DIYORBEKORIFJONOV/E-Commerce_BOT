package minao1

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"

	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/config"
	"github.com/minio/minio-go/v7"
)

type FileStorage struct {
	cfg          *config.Config
	minio_client *minio.Client
}

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
	// Открываем файл
	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("cannot open file: %s", err.Error())
	}
	defer f.Close()

	// Чтение содержимого файла в буфер
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(f); err != nil {
		return "", fmt.Errorf("failed to read file: %s", err.Error())
	}

	// Декодируем изображение для проверки формата и размера
	img, format, err := image.Decode(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %s", err.Error())
	}
	log.Println("Image format:", format)

	// Проверка размера
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	const standardWidth = 1024
	const standardHeight = 1024

	if width != standardWidth || height != standardHeight {
		return "", fmt.Errorf("invalid image size: expected %dx%d, got %dx%d",
			standardWidth, standardHeight, width, height)
	}

	// Определяем или подставляем Content-Type
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = http.DetectContentType(buf.Bytes())
	}
	log.Println("Detected Content-Type:", contentType)

	// Генерация уникального имени файла
	filename := fmt.Sprintf("%d", time.Now().UnixMilli())

	// Загрузка в MinIO
	reader := bytes.NewReader(buf.Bytes())
	_, err = s.minio_client.PutObject(
		context.Background(),
		"products",
		filename,
		reader,
		int64(reader.Len()),
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %s", err.Error())
	}

	return filename, nil
}

func (s *FileStorage) UploadFileFromBytes(data []byte, contentType string) (string, error) {
	filename := fmt.Sprintf("%d", time.Now().UnixMilli())
	reader := bytes.NewReader(data)

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
	_, err := s.minio_client.StatObject(context.Background(), "products", filename, minio.StatObjectOptions{})
	if err != nil {
		return fmt.Errorf("file not found: %s", err.Error())
	}

	err = s.minio_client.RemoveObject(context.Background(), "products", filename, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file: %s", err.Error())
	}

	return nil
}
