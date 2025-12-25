package s3

import (
	"context"
	"errors"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (s *storage) MultipartComplete(ctx context.Context, in MultipartCompleteInput) (MultipartCompleteOutput, error) {
	if in.UploadID == "" {
		return MultipartCompleteOutput{}, ErrEmptyUploadID
	}
	if len(in.Parts) == 0 {
		return MultipartCompleteOutput{}, ErrEmptyParts
	}

	parts := make([]CompletedPart, 0, len(in.Parts))
	parts = append(parts, in.Parts...)
	sort.Slice(parts, func(i, j int) bool {
		return parts[i].PartNumber < parts[j].PartNumber
	})

	s3Parts := make([]types.CompletedPart, 0, len(parts))
	for _, p := range parts {
		if p.PartNumber < 1 || p.PartNumber > 10000 {
			return MultipartCompleteOutput{}, ErrInvalidPartNumber
		}
		etag := strings.TrimSpace(p.ETag)
		if etag == "" {
			return MultipartCompleteOutput{}, errors.New("s3: empty etag")
		}

		s3Parts = append(s3Parts, types.CompletedPart{
			ETag:       &etag,
			PartNumber: &p.PartNumber,
		})
	}

	req := &s3.CompleteMultipartUploadInput{
		Bucket:   &s.bucket,
		Key:      &in.Key,
		UploadId: &in.UploadID,
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: s3Parts,
		},
	}

	out, err := s.client.CompleteMultipartUpload(ctx, req)
	if err != nil {
		return MultipartCompleteOutput{}, err
	}

	res := MultipartCompleteOutput{}
	if out.ETag != nil {
		res.ETag = *out.ETag
	}

	return res, nil

}
