package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *storage) PresignGetURL(ctx context.Context, in PresignGetInput) (PresignGetOutput, error) {
	req := &s3.GetObjectInput{
		Bucket: &s.bucket,
		Key:    &in.Key,
	}

	ps, err := s.presigner.PresignGetObject(
		ctx,
		req,
		s3.WithPresignExpires(in.TTL),
	)
	if err != nil {
		return PresignGetOutput{}, err
	}

	return PresignGetOutput{
		URL:     ps.URL,
		Headers: map[string]string{},
	}, nil
}
