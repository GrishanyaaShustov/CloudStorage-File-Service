package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) ListFolders(
	ctx context.Context,
	userID uuid.UUID,
	parentID *uuid.UUID,
) ([]domain.Folder, error) {
	rows, err := r.pool.Query(
		ctx,
		queryListFolders,
		userID,
		parentID,
	)
	if err != nil {
		return nil, mapRepoError(err)
	}
	defer rows.Close()

	folders := make([]domain.Folder, 0)

	for rows.Next() {
		var folder domain.Folder

		if err := rows.Scan(
			&folder.ID,
			&folder.UserID,
			&folder.ParentID,
			&folder.Name,
			&folder.CreatedAt,
		); err != nil {
			return nil, mapRepoError(err)
		}

		folders = append(folders, folder)
	}

	if err := rows.Err(); err != nil {
		return nil, mapRepoError(err)
	}

	return folders, nil
}
