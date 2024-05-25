package cloudcourier

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func newGcpClient(ccb *CloudCourierBridge) (StorageClient, error) {
	if ccb.CloudBucket == "" {
		return nil, fmt.Errorf("no bucket name for google cloud storage")
	}
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return &GcsClient{
		Client:     client,
		BucketName: ccb.CloudBucket,
		ctx:        ctx,
	}, nil
}

func (g *GcsClient) UploadFile(filePath string, reader io.Reader) error {
	var BaseFileName string
	ctx := context.Background()
	if filePath != "" {
		BaseFileName = filepath.Base(filePath)
	} else {
		return errors.New("you did not specify the filepath")
	}

	obj := g.Client.Bucket(g.BucketName).Object(BaseFileName)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, reader); err != nil {
		return fmt.Errorf("you did not set the reader to the file")
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("could not upload file")
	}
	v, err := obj.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println(v)
	return nil
}

type GcsClient struct {
	// The client we will use in communicating with the gcs
	Client *storage.Client
	// The name of the bucket to operate on
	BucketName string

	ctx context.Context
}

func (g *GcsClient) DeleteFile(fieldID string) error {
	return nil
}
func (g *GcsClient) ListFiles(directory string) ([]string, error) {
	// For lisiting files in a google cloud storage you have to list the nme of the bucket
	var files []string
	it := g.Client.Bucket(directory).Objects(g.ctx, nil)
	for {
		file, err := it.Next()
		if err == iterator.Done() {
			return files, nil
		}
		if err != nil {
			return nil, errors.New("")
		}
		files = append(files, file.Name)
	}
	return nil, nil
}

func (g *GcsClient) GetFile(fileID string) (io.Reader, error) {
	return nil, nil
}
