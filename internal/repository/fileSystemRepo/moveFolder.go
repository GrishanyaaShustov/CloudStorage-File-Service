package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) MoveFolder(
	ctx context.Context,
	userID uuid.UUID,
	folderID uuid.UUID,
	newParentID *uuid.UUID,
) error {
	cmd, err := r.pool.Exec(
		ctx,
		queryMoveFolder,
		userID,
		folderID,
		newParentID,
	)
	if err != nil {
		return mapRepoError(err)
	}

	if cmd.RowsAffected() == 0 {
		return domain.ErrNotFound
	}

	return nil
}
