package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/howeyc/fsnotify"
)

var (
	project  = flag.String("project", "", "The project you want to fuzz")
	bucket   = flag.String("bucket", "", "The s3 bucket you want to upload fuzz info to")
	region   = flag.String("region", "us-west-1", "The region that your s3 bucket exists in")
	dev      = flag.Bool("dev", false, "whether or not this is in dev mode.")
	S3Access = "AKIAJ4JS44QR4IIE2TMQ"
	S3Secret = "E0p2/6qd7eN95FH1NYwJHouyGA6JUYFUZq2yN5js"
	svc      *s3.S3
	sess     *session.Session
)

func handleEvents(ev *fsnotify.FileEvent) {
	// We only really care if a file is created.
	// This is because AFL should never modify
	// an existing file, or delete a file.
	// If a new file is created. Sync it up
	// to S3 with the same existing path.
	if ev.IsCreate() {
		// Create an uploader with the session and default options
		uploader := s3manager.NewUploader(sess)

		f, err := os.Open(ev.Name)
		if err != nil {
			log.Printf("Failed to open file %s: %s\n", ev.Name, err)
			return
		}

		// Upload the file to S3.
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: bucket,
			Key:    aws.String(ev.Name),
			Body:   f,
		})

		if err != nil {
			log.Printf("Could not upload file to s3: %s\n", err)
			return
		}
	}
}

func init() {
	os.Setenv("AWS_ACCESS_KEY_ID", S3Access)
	os.Setenv("AWS_SECRET_ACCESS_KEY", S3Secret)

	sess = session.Must(session.NewSession(&aws.Config{
		Region: region,
	}))
	svc = s3.New(sess)
}

func main() {
	flag.Parse()
	if *project == "" {
		log.Fatalf("Specify project name with -project")
	}

	if *bucket == "" {
		log.Fatalf("Specify bucket name with -bucket")
	}

	path := fmt.Sprintf("/fuzz/%s/out/", *project)
	if *dev == true {
		path = "out"
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				handleEvents(ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	// Wait until our watchers are in place.
	// These directories are created by AFL
	for {
		err = watcher.Watch(path + "/crashes")
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}
		break
	}

	for {
		err = watcher.Watch(path + "/hangs")
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}
		break
	}

	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()
}
