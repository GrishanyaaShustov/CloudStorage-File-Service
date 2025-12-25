package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *storage) MultipartCreate(ctx context.Context, in MultipartCreateInput) (MultipartCreateOutput, error) {
	req := &s3.CreateMultipartUploadInput{
		Bucket:      &s.bucket,
		Key:         &in.Key,
		ContentType: &in.MIMEType,
	}

	out, err := s.client.CreateMultipartUpload(ctx, req)
	if err != nil {
		return MultipartCreateOutput{}, err
	}

	if out.UploadId == nil || *out.UploadId == "" {
		return MultipartCreateOutput{}, ErrEmptyUploadID
	}

	return MultipartCreateOutput{
		UploadID: *out.UploadId,
	}, nil

}
