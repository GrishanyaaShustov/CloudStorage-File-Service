package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *storage) MultipartAbort(ctx context.Context, in MultipartAbortInput) error {
	if in.UploadID == "" {
		return ErrEmptyUploadID
	}

	req := &s3.AbortMultipartUploadInput{
		Bucket:   &s.bucket,
		Key:      &in.Key,
		UploadId: &in.UploadID,
	}

	_, err := s.client.AbortMultipartUpload(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
