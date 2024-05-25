package cloudinary_client

import "github.com/cloudinary/cloudinary-go/v2"

type Cloudinary struct {
	// Name you give to a file to be stored in cloudinary
	Tag string
	// This is for the transport side responsible for the intercommunication
	Client *cloudinary.Cloudinary
	// The Api key for now I am just using it to access the cloudinary services
	ApiKey string
	// The Api secret which corresponds to the ApiKey
	ApiSecret string
	// This is for cloudinary you need to provide the cloud name for the cloudinary
	CloudName string
}
