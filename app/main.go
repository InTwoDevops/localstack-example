package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// go run main.go -op upload my-bucket example.txt
// go run main.go -op download my-bucket example.txt
func main() {
	operation := flag.String("op", "upload", "Operation: upload or download")
	flag.Parse()

	if len(flag.Args()) != 2 {
		exitErrorf("bucket and file name required\nUsage: %s -op upload|download bucket_name filename",
			os.Args[0])
	}

	bucket := flag.Arg(0)
	filename := flag.Arg(1)

	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://localhost:4566"),
	})

	switch *operation {
	case "upload":
		uploadFile(sess, bucket, filename)
	case "download":
		downloadFile(sess, bucket, filename)
	default:
		exitErrorf("Invalid operation: %s. Operation must be 'upload' or 'download'", *operation)
	}
}

func uploadFile(sess *session.Session, bucket, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", filename, err)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}
	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}

func downloadFile(sess *session.Session, bucket, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", filename, err)
	}

	defer file.Close()
	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", filename, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
