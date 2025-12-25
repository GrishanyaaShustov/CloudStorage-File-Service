package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *storage) PresignPutURL(ctx context.Context, in PresignPutInput) (PresignPutOutput, error) {
	req := &s3.PutObjectInput{
		Bucket:      &s.bucket,
		Key:         &in.Key,
		ContentType: &in.MIMEType,
	}

	ps, err := s.presigner.PresignPutObject(
		ctx,
		req,
		s3.WithPresignExpires(in.TTL),
	)

	if err != nil {
		return PresignPutOutput{}, err
	}

	return PresignPutOutput{
		URL: ps.URL,
		Headers: map[string]string{
			"Content-Type": in.MIMEType,
		},
	}, nil
}
