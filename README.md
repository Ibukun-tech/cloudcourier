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
go get -u github.com/Ibukun-tech/cloudcourier
```

### Usage

Here's a quick example of how to use Cloud Courier to upload a file to Google cloud storage, list files from google cloud storage:

```go
package main

import (
    "github.com/Ibukun-tech/cloudcourier"
    "os"
    "log"
    "fmt"
)

func main() {
    gcp:=&cloudcourier.GcpCloud{
        Bucket:"name-Bucket"
    }

    client, err:= cloudcourier.NewCloudCourierBridge(gcp)
    if err!= nil{
        log.Fatal(err)
    }
    writer, err:= os.Open("./pathToTheFile")
    if err!= nil{
        // You can also find a better way to handle the err
        log.Fatal(err)
    }
   // To upload file in google cloud storage
   // The path is specified to get the object name where it
   // would  be storeg in Google cloud storage
   err:= client.UploadFile("/atata/ayayay", writer)
    if err!= nil{
        log.Fatal(err)
    }

    // To also list files
    // You put the name of the bucket you want to
    files, err:=client.ListFiles("name-bucket")
    if err!= nil{
        log.Fatal(err)
    }
    for _,v:=range files{
        fmt.Println(v)
    }
}
```

### Usage of cloudcourier with AWS

```go
import(
    "github.com/Ibukun-tech/cloudcourier"
    "log"
    "fmt"
    "io"
)

func main(){
    aws:=&cloudcourier.AwsManufacture{
        Region:"aws-region",
        Bucket:"aws-bucket",
    }
    client, err:=cloudcourier.NewCloudCourierBridge(aws)
    if err!= nil{
        log.Fatal(err)
    }
    writer, err:= os.Open("./pathToTheFile")
    if err!= nil{
        // You can also find a better way to handle the err
        log.Fatal(err)
    }
    // This how you upload file to AWS
    err:= client.UploadFile("/atata/ayayay", writer)
    if err!= nil{
        log.Fatal(err)
    }

    // DeleteFile deletes a file from S3 by its key.
    err:=client.DeleteFile("key")
    if err!= nil{
        log.Fatal(err)
    }
    // ListFiles lists all files in a specified directory of the   S3 bucket.
    files, err:=client.ListFiles("directory")
    if err!= nil{
        log.Fatal(err)
    }
    // GetFile retrieves a file as an io.Reader by its key from S3.
    var body io.Reader
    var err error
    body, err=client.GetFile("key")
    if err!= nil{
        log.Fatal(err)
    }

}
```

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
