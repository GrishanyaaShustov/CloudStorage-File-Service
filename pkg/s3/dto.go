package s3

import "time"

type PresignPutInput struct {
	Bucket    string
	Key       string
	TTL       time.Duration
	MIMEType  string
	SizeBytes int64
}

type PresignPutOutput struct {
	URL     string
	Headers map[string]string
}

type MultipartCreateInput struct {
	Bucket   string
	Key      string
	TTL      time.Duration
	MIMEType string
}

type MultipartCreateOutput struct {
	UploadID string
}

type MultipartPresignPartInput struct {
	Bucket     string
	Key        string
	UploadID   string
	PartNumber int32
	TTL        time.Duration
}

type MultipartPresignPartOutput struct {
	URL     string
	Headers map[string]string
}

type CompletedPart struct {
	PartNumber int32
	ETag       string
}

type MultipartCompleteInput struct {
	Bucket   string
	Key      string
	UploadID string
	Parts    []CompletedPart // must be sorted by PartNumber (S3 requirement)
}

type MultipartCompleteOutput struct {
	ETag string // ETag final object (optional)
}

type MultipartAbortInput struct {
	Bucket   string
	Key      string
	UploadID string
}

type DeleteObjectInput struct {
	Bucket string
	Key    string
}

type PresignGetInput struct {
	Bucket string
	Key    string
	TTL    time.Duration
}

type PresignGetOutput struct {
	URL     string
	Headers map[string]string
}
