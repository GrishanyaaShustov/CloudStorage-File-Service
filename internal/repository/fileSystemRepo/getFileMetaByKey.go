package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) GetFileMetaByKey(
	ctx context.Context,
	userID uuid.UUID,
	s3Key string,
) (domain.File, error) {
	var file domain.File

	err := r.pool.QueryRow(
		ctx,
		queryGetFileMetaByKey,
		userID,
		s3Key,
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
