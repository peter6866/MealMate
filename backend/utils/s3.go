package utils

import (
	"bytes"
	"fmt"
	"image"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/disintegration/imaging"
	"github.com/peter6866/foodie/config"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func UploadFileToS3(file *multipart.FileHeader, crop bool) (string, error) {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Decode the image
	img, err := imaging.Decode(src, imaging.AutoOrientation(true))
	if err != nil {
		return "", err
	}

	var resizedImg *image.NRGBA

	if crop {
		size := minInt(img.Bounds().Dx(), img.Bounds().Dy())

		// resize and crop to a square
		resizedImg = imaging.Fill(img, size, size, imaging.Center, imaging.Lanczos)
	} else {
		resizedImg = imaging.Clone(img)
	}

	// create a buffer to store the resized image
	buf := new(bytes.Buffer)
	err = imaging.Encode(buf, resizedImg, imaging.JPEG, imaging.JPEGQuality(70))
	if err != nil {
		return "", err
	}

	// create a unique file name
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ".jpg")

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
		Body:                 bytes.NewReader(buf.Bytes()),
		ContentLength:        aws.Int64(int64(buf.Len())),
		ContentType:          aws.String("image/jpeg"),
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
