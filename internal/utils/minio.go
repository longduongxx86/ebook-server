package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"path/filepath"
	"strings"
	"time"

	"main/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client
var minioBucketName string
var minioEndpoint string

// InitMinIO initializes MinIO client
func InitMinIO() error {
	endpoint := config.GetEnv("MINIO_ENDPOINT", "localhost:9000")
	accessKeyID := config.GetEnv("MINIO_ACCESS_KEY", "minioadmin")
	secretAccessKey := config.GetEnv("MINIO_SECRET_KEY", "minioadmin")
	useSSL := config.GetEnv("MINIO_USE_SSL", "false") == "true"
	bucketName := config.GetEnv("MINIO_BUCKET", "ebooks")

	minioEndpoint = endpoint
	minioBucketName = bucketName

	// Initialize minio client object
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize MinIO client: %w", err)
	}

	minioClient = client

	// Check if bucket exists, if not create it
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to check bucket existence: %w", err)
	}

	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create bucket: %w", err)
		}
	}

	// Set bucket policy to public read
	policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {"AWS": ["*"]},
				"Action": ["s3:GetObject"],
				"Resource": ["arn:aws:s3:::%s/*"]
			}
		]
	}`, bucketName)

	if err := client.SetBucketPolicy(ctx, bucketName, policy); err != nil {
		log.Printf("Warning: failed to set bucket policy: %v", err)
		// Don't return error here, as MinIO might be running in gateway mode or other restrictions
	}

	return nil
}

// GetMinIOClient returns the MinIO client instance
func GetMinIOClient() *minio.Client {
	return minioClient
}

// UploadFile uploads a file to MinIO and returns the file URL
func UploadFile(file io.Reader, fileName string, contentType string, folder string) (string, error) {
	if minioClient == nil {
		return "", fmt.Errorf("MinIO client not initialized, call InitMinIO() first")
	}

	// Generate unique filename with timestamp to avoid conflicts
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(filepath.Base(fileName), ext)
	timestamp := time.Now().Unix()
	uniqueFileName := fmt.Sprintf("%s_%d%s", baseName, timestamp, ext)

	// Build object path with folder
	objectPath := uniqueFileName
	if folder != "" {
		objectPath = strings.Trim(folder, "/") + "/" + uniqueFileName
	}

	// Detect content type if not provided
	if contentType == "" {
		contentType = mime.TypeByExtension(ext)
		if contentType == "" {
			contentType = "application/octet-stream"
		}
	}

	// Upload file
	ctx := context.Background()
	_, err := minioClient.PutObject(ctx, minioBucketName, objectPath, file, -1, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// Return file URL
	fileURL := GetFileURL(objectPath)
	return fileURL, nil
}

// UploadFileBytes uploads file from bytes to MinIO
func UploadFileBytes(fileData []byte, fileName string, contentType string, folder string) (string, error) {
	reader := bytes.NewReader(fileData)
	return UploadFile(reader, fileName, contentType, folder)
}

// DeleteFile deletes a file from MinIO
func DeleteFile(objectPath string) error {
	if minioClient == nil {
		return fmt.Errorf("MinIO client not initialized, call InitMinIO() first")
	}

	ctx := context.Background()
	err := minioClient.RemoveObject(ctx, minioBucketName, objectPath, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// GetFileURL returns the public URL for a file in MinIO
func GetFileURL(objectPath string) string {
	baseURL := config.GetEnv("MINIO_BASE_URL", "")
	if baseURL != "" {
		return strings.TrimRight(baseURL, "/") + "/" + minioBucketName + "/" + objectPath
	}

	// Fallback: construct URL from endpoint
	protocol := "http"
	if config.GetEnv("MINIO_USE_SSL", "false") == "true" {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s/%s/%s", protocol, minioEndpoint, minioBucketName, objectPath)
}

// ExtractObjectPathFromURL extracts object path from full URL
func ExtractObjectPathFromURL(url string) string {
	// Remove base URL and bucket name to get object path
	parts := strings.Split(url, minioBucketName+"/")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

// CheckFileExists checks if a file exists in MinIO
func CheckFileExists(objectPath string) (bool, error) {
	if minioClient == nil {
		return false, fmt.Errorf("MinIO client not initialized, call InitMinIO() first")
	}

	ctx := context.Background()
	_, err := minioClient.StatObject(ctx, minioBucketName, objectPath, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
