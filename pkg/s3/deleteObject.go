package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *storage) DeleteObject(ctx context.Context, in DeleteObjectInput) error {
	req := &s3.DeleteObjectInput{
		Bucket: &s.bucket,
		Key:    &in.Key,
	}

	_, err := s.client.DeleteObject(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
