package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) GetFileMeta(
	ctx context.Context,
	userID uuid.UUID,
	fileID uuid.UUID,
) (domain.File, error) {
	var file domain.File

	err := r.pool.QueryRow(
		ctx,
		queryGetFileMeta,
		userID,
		fileID,
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
