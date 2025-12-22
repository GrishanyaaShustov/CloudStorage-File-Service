package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) CreateFileMeta(
	ctx context.Context,
	userID uuid.UUID,
	folderID *uuid.UUID,
	fileName string,
	s3Bucket string,
	s3Key string,
	mimeType string,
	size int64,
) (domain.File, error) {
	var file domain.File

	err := r.pool.QueryRow(
		ctx,
		queryCreateFileMeta,
		userID,
		folderID,
		fileName,
		s3Bucket,
		s3Key,
		mimeType,
		size,
	).Scan(
		&file.ID,
		&file.UserID,
		&file.FolderID,
		&file.Name,
		&file.S3Bucket,
		&file.S3Key,
		&file.MIMEType,
		&file.SizeBytes,
		&file.Status,
		&file.CreatedAt,
		&file.UpdatedAt,
	)

	if err != nil {
		return domain.File{}, mapRepoError(err)
	}

	return file, nil
}
