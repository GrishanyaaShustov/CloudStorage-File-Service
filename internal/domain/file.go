package domain

import (
	"time"

	"github.com/google/uuid"
)

type FileStatus string

const (
	FileStatusUploading FileStatus = "uploading"
	FileStatusUploaded  FileStatus = "uploaded"
)

type File struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	FolderID *uuid.UUID // nil = корень

	Name     string
	S3Bucket string
	S3Key    string
	MIMEType string

	SizeBytes int64

	Status    FileStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
