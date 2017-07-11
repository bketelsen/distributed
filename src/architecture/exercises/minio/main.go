package main

import (
	"log"

	minio "github.com/minio/minio-go"
)

/* Get your key and id from docker logs
Endpoint:  http://172.17.0.2:9000  http://127.0.0.1:9000
AccessKey: HID6B4S9ANGIGGW435G3
SecretKey: 9HlelpHZNnjmXgLBB/CBofcuQUepbEX2sf0iv3DW
Region:    us-east-1
SQS ARNs:  <none>
*/

func main() {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "81SWL1ZIVK596J52R0ID"
	secretAccessKey := "ro0GIanW/HNHr1vyjKiWbX09Ms3zpu1BAwQSUD2G"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucked
	bucketName := "gophertrain"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	// Upload the zip file
	objectName := "main.go"
	filePath := "./main.go"
	contentType := "text/plain"

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, contentType)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
