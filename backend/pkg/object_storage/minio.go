package object_storage

import (
	"buddylink/config"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"net/http"
	"time"
)

type MinioClient interface {
	UploadFileToBucket(bucket string, filename string, reader io.Reader) error
	GetUrl(bucket string, filename string) (string, error)
	DeleteFileFormBucket(bucket, filename string) error
}

type MinioImpl struct {
	client *minio.Client
	ctx    context.Context
}

func NewMinio(cfg config.MinioConfig) (MinioClient, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	return &MinioImpl{
		client: client,
		ctx:    context.Background(),
	}, nil
}

func (m *MinioImpl) UploadFileToBucket(bucket string, filename string, reader io.Reader) error {

	// 检查桶是否存在
	exists, err := m.client.BucketExists(m.ctx, bucket)
	if err != nil {
		log.Println("Error checking if bucket exists: ", err)
		return err
	}

	// 桶不存在就创建
	if !exists {
		if err := m.client.MakeBucket(m.ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			return fmt.Errorf("Error creating bucket %s: %s", bucket, err)
		}
		log.Println("Successfully created bucket: ", bucket)
	}

	// 上传文件
	_, err = m.client.PutObject(m.ctx, bucket, filename, reader, -1, minio.PutObjectOptions{})

	if err != nil {
		log.Println("Error uploading file: ", err)
		return err
	}
	log.Println("Successfully uploaded file: ", filename)
	return nil
}

func (m *MinioImpl) GetUrl(bucket string, filename string) (string, error) {
	/*
		获取文件访问的URL
	*/

	exists, err := m.client.BucketExists(m.ctx, bucket)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("bucket does not exist")
	}

	// 检查文件是否存在
	_, err = m.client.StatObject(m.ctx, bucket, filename, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).StatusCode == http.StatusNotFound {
			return "", fmt.Errorf("file not found")
		}
		log.Printf("Error getting file %s: %s", filename, err)
		return "", err
	}

	url, err := m.client.PresignedGetObject(m.ctx, bucket, filename, time.Hour*24*365*10, nil)
	if err != nil {
		log.Printf("Error getting file %s: %s", filename, err)
		return "", err
	}
	return url.String(), nil
}

func (m *MinioImpl) DeleteFileFormBucket(bucket, filename string) error {
	exists, err := m.client.BucketExists(m.ctx, bucket)
	if err != nil {
		log.Println("Error checking if bucket exists: ", err)
		return err
	}
	if !exists {
		log.Println("Bucket does not exist")
		return fmt.Errorf("Bucket does not exist")
	}

	_, err = m.client.StatObject(m.ctx, bucket, filename, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).StatusCode == http.StatusNotFound {
			log.Println("file does not exist")
			return fmt.Errorf("file does not exist")
		}
		log.Println("Error checking if file exists: ", err)
		return err
	}

	if err := m.client.RemoveObject(m.ctx, bucket, filename, minio.RemoveObjectOptions{}); err != nil {
		return err
	}
	log.Println("Successfully removed file: ", filename)
	return nil
}
