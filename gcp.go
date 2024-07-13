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

type GcpCloud struct {
	Bucket string
}

func (g *GcpCloud) GetProvider() cloudCourierProvider {
	return GCP
}

func init() {
	storeFunc[GCP] = newGcpClient
}

type GcsClient struct {
	// The client we will use in communicating with the gcs
	Client *storage.Client
	// The name of the bucket to operate on
	BucketName string
	ctx        context.Context
}

func newGcpClient(ccb Provider) (StorageClient, error) {
	gcConfig, ok := ccb.(*GcpCloud)
	// if ccb.type
	if !ok {
		return nil, fmt.Errorf("incorrect configuration")
	}
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return &GcsClient{
		Client:     client,
		BucketName: gcConfig.Bucket,
		ctx:        ctx,
	}, nil
}

func (g *GcsClient) UploadFile(filePath string, reader io.Reader) error {
	var BaseFileName string
	if filePath != "" {
		BaseFileName = filepath.Base(filePath)
	} else {
		return errors.New("you did not specify the filepath")
	}

	obj := g.Client.Bucket(g.BucketName).Object(BaseFileName)
	w := obj.NewWriter(g.ctx)
	if _, err := io.Copy(w, reader); err != nil {
		return fmt.Errorf("you did not set the reader to the file")
	}
	defer func() error {
		if err := w.Close(); err != nil {
			return fmt.Errorf("could not upload file")
		}
		return nil
	}()

	v, err := obj.Attrs(g.ctx)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println(v)
	return nil
}

func (g *GcsClient) ListFiles(directory string) ([]string, error) {
	// For lisiting files in a google cloud storage you have to list the name of the bucket
	var files []string
	it := g.Client.Bucket(directory).Objects(g.ctx, nil)
	for {
		file, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.New("")
		}
		files = append(files, file.Name)
	}
	return files, nil
}

func (g *GcsClient) GetFile(fileID string) (io.Reader, error) {
	return nil, nil
}

func (g *GcsClient) DeleteFile(fieldID string) error {
	return nil
}
