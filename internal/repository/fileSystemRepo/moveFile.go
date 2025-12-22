package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) MoveFile(
	ctx context.Context,
	userID uuid.UUID,
	fileID uuid.UUID,
	newFolderID *uuid.UUID,
) error {
	cmd, err := r.pool.Exec(
		ctx,
		queryMoveFile,
		userID,
		fileID,
		newFolderID,
	)
	if err != nil {
		return mapRepoError(err)
	}

	if cmd.RowsAffected() == 0 {
		// файл не найден или не принадлежит пользователю
		return domain.ErrNotFound
	}

	return nil
}
