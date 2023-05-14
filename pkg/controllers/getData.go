package controller

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/samridht23/safepasteServer/pkg/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetData(c *gin.Context) {
	bucketRegion := utils.Getenv("AWS_BUCKET_REGION")
	bucketName := utils.Getenv("AWS_BUCKET_NAME")
	accessID := utils.Getenv("AWS_ACCESS_ID")
	secretKey := utils.Getenv("AWS_SECRET_KEY")
	file_key := c.Request.URL.Query()["file_key"][0]
	var encryptedString string = ""
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(bucketRegion),
		Credentials: credentials.NewStaticCredentials(accessID, secretKey, ""),
	}))
	svc := s3.New(sess)
	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file_key),
	}
	resp, err := svc.GetObject(getObjectInput)
	if err != nil {
		log.Fatal("Error getting Object", err)
	}
	if resp.ServerSideEncryption != nil {
		switch *resp.ServerSideEncryption {
		case "AES256":
			plaintext, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading object:", err)
				os.Exit(1)
			}
			encryptedString = string(plaintext)
		default:
			fmt.Println("Unsupported encryption algorithm:", *resp.ServerSideEncryption)
			os.Exit(1)
		}
	} else {
		// Object is not encrypted, so just read the contents
		plaintext, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading object:", err)
			os.Exit(1)
		}
		fmt.Println("Object contents:", string(plaintext))
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    fmt.Println(encryptedString)
	c.JSON(http.StatusOK, gin.H{"encryptedString": encryptedString})
}
