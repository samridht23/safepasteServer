package controller

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/samridht23/safepasteServer/pkg/utils"
	"log"
	"fmt"
	"net/http"
)

func UploadData(c *gin.Context) {
	bucketRegion := utils.Getenv("AWS_BUCKET_REGION")
	bucketName := utils.Getenv("AWS_BUCKET_NAME")
	accessID := utils.Getenv("AWS_ACCESS_ID")
	secretKey := utils.Getenv("AWS_SECRET_KEY")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(bucketRegion),
		Credentials: credentials.NewStaticCredentials(accessID, secretKey, ""),
	}))
	svc := s3.New(sess)
	var requestBody struct {
		Expiry          float64 `json:"expiry"`
		EncryptedString string  `json:"encryptedString"`
		UniqueId        string  `json:"uniqueId"`
	}
	err := c.BindJSON(&requestBody)
	objectContent := []byte(requestBody.EncryptedString)
    fmt.Println("Encrypted objectContent from the server and from upload request",objectContent)
	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(requestBody.UniqueId),
		Body:   bytes.NewReader(objectContent),
	}
	resp, err := svc.PutObject(putObjectInput)
	log.Print("Upload data response object", resp)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Uploaded"})
}
