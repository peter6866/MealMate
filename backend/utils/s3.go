package utils

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/peter6866/foodie/config"
)

func UploadFileToS3(file *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Read the content of the uploaded file
	size := file.Size
	buffer := make([]byte, size)
	_, err = src.Read(buffer)
	if err != nil {
		return "", err
	}

	// create a unique file name
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))

	// configure aws session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AppConfig.AWS_REGION),
		Credentials: credentials.NewStaticCredentials(
			config.AppConfig.AWS_ACCESS_KEY,
			config.AppConfig.AWS_SECRET_ACCESS_KEY,
			"",
		),
	})
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(config.AppConfig.AWS_S3_BUCKET),
		Key:                  aws.String(filename),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		return "", err
	}

	// get the url of the uploaded file
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", config.AppConfig.AWS_S3_BUCKET, config.AppConfig.AWS_REGION, filename)

	return url, nil
}

func DeleteFileFromS3(fileURL string) error {
	// Extract the key (filename) from the URL
	u, err := url.Parse(fileURL)
	if err != nil {
		return err
	}
	key := strings.TrimPrefix(u.Path, "/")

	// Configure AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AppConfig.AWS_REGION),
		Credentials: credentials.NewStaticCredentials(
			config.AppConfig.AWS_ACCESS_KEY,
			config.AppConfig.AWS_SECRET_ACCESS_KEY,
			"",
		),
	})
	if err != nil {
		return err
	}

	// Create S3 service client
	svc := s3.New(sess)

	// Delete the object
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(config.AppConfig.AWS_S3_BUCKET),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return nil
}
