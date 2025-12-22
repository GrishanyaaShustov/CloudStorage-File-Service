package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) ListFilesMeta(
	ctx context.Context,
	userID uuid.UUID,
	folderID *uuid.UUID,
) ([]domain.File, error) {
	rows, err := r.pool.Query(
		ctx,
		queryListFilesMeta,
		userID,
		folderID,
	)
	if err != nil {
		return nil, mapRepoError(err)
	}
	defer rows.Close()

	files := make([]domain.File, 0)

	for rows.Next() {
		var file domain.File

		if err := rows.Scan(
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
		); err != nil {
			return nil, mapRepoError(err)
		}

		files = append(files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, mapRepoError(err)
	}

	return files, nil
}
