package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *storage) MultipartPresignPartURL(ctx context.Context, in MultipartPresignPartInput) (MultipartPresignPartOutput, error) {
	if in.UploadID == "" {
		return MultipartPresignPartOutput{}, ErrEmptyUploadID
	}
	if in.PartNumber < 1 || in.PartNumber > 10000 {
		return MultipartPresignPartOutput{}, ErrInvalidPartNumber
	}

	req := &s3.UploadPartInput{
		Bucket:     &s.bucket,
		Key:        &in.Key,
		UploadId:   &in.UploadID,
		PartNumber: &in.PartNumber,
	}

	ps, err := s.presigner.PresignUploadPart(
		ctx,
		req,
		s3.WithPresignExpires(in.TTL),
	)
	if err != nil {
		return MultipartPresignPartOutput{}, err
	}

	return MultipartPresignPartOutput{
		URL:     ps.URL,
		Headers: map[string]string{},
	}, nil

}
