package fileSystemRepo

import (
	"context"

	"github.com/google/uuid"
)

func (r *repo) FolderExists(
	ctx context.Context,
	userID uuid.UUID,
	folderID uuid.UUID,
) (bool, error) {
	var exists bool

	err := r.pool.QueryRow(
		ctx,
		queryFolderExists,
		userID,
		folderID,
	).Scan(&exists)
	if err != nil {
		return false, mapRepoError(err)
	}

	return exists, nil
}
