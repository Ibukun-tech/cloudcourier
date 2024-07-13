# Cloud Courier

Cloud Courier is a Go library designed to simplify the process of uploading files to various cloud storage services.
It provides a unified interface to interact with different cloud providers, allowing developers to integrate cloud storage capabilities into their applications seamlessly.

## Features

- **Unified API**: Use a single, consistent API to interact with multiple cloud storage providers.
- **Extensible**: Easily add support for additional cloud providers.
- **Simple Configuration**: Configure your cloud storage credentials and settings in one place.
- **Error Handling**: Robust error handling for common edge cases in file uploads.

## Supported Cloud Providers

- Google cloud storage
- Aws Bucket

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

To install Cloud Courier, use the following `go get` command:

```sh
go get https://github.com/Ibukun-tech/cloudcourier
```

### Usage

Here's a quick example of how to use Cloud Courier to upload a file to Cloudinary:

```go
package main

import (
    "github.com/Ibukun-tech/cloudcourier"
)

func main() {
    // Initialize the CloudCourierBridge with your Cloudinary credentials
    ccb := &types.CloudCourierBridge{
        CloudProvider: "cloudinary",
        ApiKey:        "your-api-key",
        ApiSecret:     "your-api-secret",
        CloudName:     "your-cloud-name",
    }

    // Create a new Cloud Courier instance
    courier, err := cloudcourier.NewCloudCourier(ccb)
    if err != nil {
        log.Fatalf("Failed to create Cloud Courier: %v", err)
    }

    // Upload a file
    err = courier.UploadFile(yourFile)
    if err != nil {
        log.Fatalf("Failed to upload file: %v", err)
    }
}
```

Replace `your-api-key`, `your-api-secret`, and `your-cloud-name` with your actual Cloudinary credentials.

### Documentation

For detailed documentation, refer to the `docs` directory in this repository.

## Contributing

We welcome contributions to Cloud Courier\! Please read our [CONTRIBUTING.md](CONTRIBUTING.md) file to see how you can help improve this project.

## License

Cloud Courier is released under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Acknowledgments

- Thanks to all the contributors who have helped with the development of Cloud Courier.

## Contact

For questions and feedback, please reach out to the maintainers at <oyetunjiibukunoluwa8@gmail.com>.
