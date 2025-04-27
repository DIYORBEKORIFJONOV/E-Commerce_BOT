package minao

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)


type Client struct {
    mc         *minio.Client
    bucketName string
}


func NewClient(endpoint, accessKey, secretKey, bucketName string, useSSL bool) (*Client, error) {
    mc, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        return nil, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    exists, err := mc.BucketExists(ctx, bucketName)
    if err != nil {
        return nil, err
    }
    if !exists {
        if err := mc.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
            return nil, err
        }
        log.Println("Created bucket", bucketName)
    }

    return &Client{mc: mc, bucketName: bucketName}, nil
}
func (c *Client) AddPhoto(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
    ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
    if ext != ".png" && ext != ".jpg" {
        return "", fmt.Errorf("only .png or .jpg files are allowed")
    }

    src, err := fileHeader.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    objectName := uuid.New().String() + ext

    _, err = c.mc.PutObject(ctx,
        c.bucketName,
        objectName,
        src,
        fileHeader.Size,
        minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")},
    )
    if err != nil {
        return "", err
    }

    return objectName, nil
}

func (c *Client) GetPhoto(ctx context.Context, objectName string) (io.ReadCloser, error) {
    obj, err := c.mc.GetObject(ctx, c.bucketName, objectName, minio.GetObjectOptions{})
    if err != nil {
        return nil, err
    }
    return obj, nil
}
