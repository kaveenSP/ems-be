package amazon_s3

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func InitiateConnectionWithImageService() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials("AKIA4OYELSC3GDKLEZ4X", "xIHKldKnd8PAV7VpVwfAefHt19W6eoWFAkukhdK2", ""),
	})
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)
	return svc, nil
}

func UploadImageToS3(svc *s3.S3, imagePath string) (string, error) {
	bucketName := "ems-planners-images"
	imageKey := "images/" + uuid.New().String()
	imageData, err := base64.StdEncoding.DecodeString(imagePath)
	if err != nil {
		return "", err
	}
	imageReader := bytes.NewReader(imageData)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(imageKey),
		Body:   imageReader,
	})
	if err != nil {
		return "", err
	}
	return "https://s3.ap-south-1.amazonaws.com/" + bucketName + "/" + imageKey, nil
}
