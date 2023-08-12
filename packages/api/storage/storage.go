package storage

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func getMinioClient() (*minio.Client, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := os.Getenv("MINIO_SSL") == "true"

	if accessKey == "" || secretKey == "" || endpoint == "" {
		return nil, fmt.Errorf("MinIO credentials not set")
	}

	// Initialize MinIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to initialize MinIO client: %v", err)
	}

	return minioClient, nil
}

func GenerateUploadURL(c *gin.Context) {
	bucketName := os.Getenv("MINIO_BUCKET")
	expireTimeStr := os.Getenv("MINIO_UPLOAD_EXPIRE_TIME")

	expireTime, err := strconv.Atoi(expireTimeStr)
	if err != nil || expireTime <= 0 {
		expireTime = 1000
	}

	minioClient, err := getMinioClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var payload struct {
		UploadKey string `json:"uploadName"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	objectName := "uploads/" + payload.UploadKey
	uploadURL, err := minioClient.PresignedPutObject(context.Background(), bucketName, objectName, time.Duration(expireTime)*time.Second)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate upload token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"upload_url": uploadURL})
}

func ServePrivateObject(c *gin.Context) {
	bucketName := os.Getenv("MINIO_BUCKET")
	expireTimeStr := os.Getenv("MINIO_UPLOAD_EXPIRE_TIME")

	expireTime, err := strconv.Atoi(expireTimeStr)
	if err != nil || expireTime <= 0 {
		expireTime = 1000
	}

	minioClient, err := getMinioClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	objectKey := c.Param("key")
	objectKey = strings.ReplaceAll(objectKey, "..", "")
	objectKey = strings.TrimPrefix(objectKey, "/")

	// Set request parameters
	filename := path.Base(objectKey)
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Generate a presigned URL to access the private image
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectKey, time.Duration(expireTime)*time.Second, reqParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate presigned URL"})
		return
	}

	// Redirect to the presigned URL to serve the private image
	c.Redirect(http.StatusTemporaryRedirect, presignedURL.String())
}
