# Preliminaries

- who are your users? are they developers? yes.
- why would they use your SDK? to easily switch bridges to different cloud providers when uploading files like: user profile images, audio files, video files, and others.
- how can we support these files? use golang's multipart file type.

## Code structure and organization

Since you're creating an open source SDK to simplify a process for other people, you have to use a
proper code structure that is easy to understand. Even more important is the fact that your code-based imports and exports
of identifiers needs to be sustainable.

Your project structure is currently like this:

```shell
.
/__ File
│   └── struct.go
├── README.md
├── cloudcourier.go
├── cloudinary
│   ├── cld.go
│   └── struct.go
├── cmd
│   └── api
│       └── main.go
├── config
│   └── services.go
├── const.go
├── go.mod
├── go.sum
├── helpCloudcourier.go
├── ibk.txt
├── reviews.md
├── struct.go
└── types
    └── struct.go

```

It is not neat. `File` directory should solely focus on file handling. The `struct.go` inside should be renamed to something else.
Don't use standard Golang keywords for your own implementations. At least, `file_struct.go` is better. I generally use an `internal/models.go`
and store all my generic models there. then a `<package_name>/model.go` to store models pertaining to a particular sub-directory.
Same thing for `const`, its a Go keyword, so no `const.go`.

Each cloud provider (Cloudinary, AWS S3) should have its own directory similar to the way you have `cloudinary`. That way, its
easier to extend the SDK with additional providers.

Your `config/services.go` and `const.go` all contain constants. Store all in one place, or use your config for the right reason.
It should contain configurations that your user should provide before being able to use `cloudcourier`, like their own Cloudinary API key, AWS credentials, etc.

You can refactor the codebase to this:

```shell
.
├── README.md // I already added this in my branch. It could have helped me understand the codebase better, if you added it.
├── cloudcourier.go
├── cloud_providers
│   ├── cloudinary
│   │   ├── cloudinary.go  // Renamed from cld.go, includes provider-specific implementations
│   │   └── struct.go
│   ├── aws_s3
│   │   ├── s3.go
│   │   └── struct.go
│   └── interface.go       // General interface for all providers
├── cmd
│   └── api
│       └── main.go
├── config
│   └── config.go          // Includes all configuration-related constants and functions
├── file_handling
│   └── file.go            // Merged and renamed from File/struct.go
├── go.mod
├── go.sum
├── utils
│   └── helpers.go         // Optional, for helper functions used across the project
└── reviews.md
```

## Commenting

Commenting code for an SDK or anything that Open Source users will consume is
very important. So always add comments to identifiers you create. And commenting
in Go is not just anyhow. It has a pattern you have to adhere to.
I've created an example file.go that practicalizes these suggestions.

## Returns

Always return an `error` type in every function possible, as long as its open source.
Its a standard practice.

## Naming convention

`General` is not a good name. Give clearer names to identifiers, so that your SDK users can understand what each of them does.
See an example in

## Flow

You should handle things sequentially. First, start the file management.
After that, the cloud provider management.
After that, the cloudcourier bridge to the cloud providers.
After that, the file upload to cloudcourier.
After that, the transfer of file from cloudcourier bridge to the provider.
After that, cleanups and review of flow.

## Files

- a single file can be saved.
- multiple files can be saved.
- an entire directory of files can be saved.
- create file management in its own sub-directory (called package in Go)
- file object/struct must have specific
