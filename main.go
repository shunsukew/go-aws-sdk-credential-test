package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func main() {
	region := os.Getenv("REGION")
	keyID := os.Getenv("KEY_ID")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	svc := kms.New(sess)

	algorithm := kms.SigningAlgorithmSpecEcdsaSha256
	text := "1234567890"

	// Sign the data
	result, err := svc.Sign(&kms.SignInput{
		KeyId:            aws.String(keyID),
		Message:          []byte(text),
		SigningAlgorithm: &algorithm,
	})

	if err != nil {
		fmt.Println("Got error encrypting data: ", err)
		os.Exit(1)
	}

	fmt.Println("Signature")
	fmt.Println(result.Signature)
}
