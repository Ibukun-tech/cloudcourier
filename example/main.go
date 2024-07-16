package main

import (
	"fmt"
	"log"

	"github.com/Ibukun-tech/cloudcourier"
)

func main() {
	var err error
	var client cloudcourier.StorageClient
	aws := &cloudcourier.AwsManufacture{
		Region: "aws-region",
		Bucket: "aws-bucket",
	}
	client, err = cloudcourier.NewCloudCourierBridge(aws)
	if err != nil {
		log.Fatal(err)
	}

	w, err := client.GetFile("key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(w)
}
