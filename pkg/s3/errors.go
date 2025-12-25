package s3

import "errors"

var (
	ErrEmptyUploadID     = errors.New("s3: empty upload id")
	ErrInvalidPartNumber = errors.New("s3: invalid part number (must be 1..10000)")
	ErrEmptyParts        = errors.New("s3: empty parts")
)
