package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage interface {
	// PresignPutURL returns a presigned URL for a single PUT upload.
	// Intended for objects <= 5GB.
	PresignPutURL(ctx context.Context, in PresignPutInput) (PresignPutOutput, error)

	// MultipartCreate starts a multipart upload and returns upload ID.
	MultipartCreate(ctx context.Context, in MultipartCreateInput) (MultipartCreateOutput, error)

	// MultipartPresignPartURL returns a presigned URL to upload a single part.
	// partNumber must be in [1..10000].
	MultipartPresignPartURL(ctx context.Context, in MultipartPresignPartInput) (MultipartPresignPartOutput, error)

	// MultipartComplete completes multipart upload with the list of uploaded parts (ETags).
	MultipartComplete(ctx context.Context, in MultipartCompleteInput) (MultipartCompleteOutput, error)

	// MultipartAbort aborts multipart upload (cleanup).
	MultipartAbort(ctx context.Context, in MultipartAbortInput) error

	// DeleteObject deletes an object by bucket/key.
	DeleteObject(ctx context.Context, in DeleteObjectInput) error

	// PresignGetURL returns a presigned URL for download.
	PresignGetURL(ctx context.Context, in PresignGetInput) (PresignGetOutput, error)
}

type storage struct {
	client    *s3.Client
	presigner *s3.PresignClient
	bucket    string
}

func New(client *s3.Client, presigner *s3.PresignClient, bucket string) Storage {
	return &storage{
		client:    client,
		presigner: presigner,
		bucket:    bucket,
	}
}
