package cloudinary_client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	file "github.com/Ibukun-tech/cloudcourier/File"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

//	UploadFile(filePath string, reader io.Reader) error // Uploads a file from an io.Reader source to the specified path.
//
// DeleteFile(fileID string) error                     // Deletes a file identified by a unique identifier.
// ListFiles(directory string) ([]string, error)       // Lists files under a specified directory.
// GetFile(fileID string) (io.Reader, error)           // Retrieves a file as an io.Reader by its unique identifier.

func (c *Cloudinary) UploadFile(fileInterface interface{}) error {
	fmt.Println("i AM HERE WORKING ALREADY ------ 1")
	files, ok := fileInterface.(file.File)
	if !ok {
		return fmt.Errorf("it must be of the type cloudcourier.File")
	}
	// To check if its if the file fits the requirement
	if err := files.CheckIfTheFileIsValid(); err != nil {
		return err
	}

	// var err error
	if files.Path != "" {

		_, err := os.Stat(files.Path)

		if err != nil {
			return fmt.Errorf("this path does not exist")
		}
		if os.IsNotExist(err) {
			return fmt.Errorf("this path does not exist")
		}
		file, err := os.Open(files.Path)
		if err != nil {
			return fmt.Errorf("this file does not exist")
		}
		defer file.Close()
		files.ToHandle = file

		if files.FileName == "" {
			files.FileName = filepath.Base(files.Path)
		}
	}

	ctx := context.Background()
	// Work on getting a random public Id or I can specify it from the
	// I feel i need to implement the side like a go routine why because I am upload the file into the storage activite and also trying to get resources from it
	resp, err := c.Client.Upload.Upload(ctx, files.FileName, uploader.UploadParams{PublicID: c.CloudName})
	if err != nil {
		return fmt.Errorf("%s here abi", err)
	}
	fmt.Println(resp)
	return nil
}

func (c *Cloudinary) DeletFile()
func (c *Cloudinary) ListFiles()
func (c *Cloudinary) GetFile()
