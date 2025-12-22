package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) RenameFile(
	ctx context.Context,
	userID uuid.UUID,
	fileID uuid.UUID,
	newName string,
) error {
	cmd, err := r.pool.Exec(
		ctx,
		queryRenameFile,
		userID,
		fileID,
		newName,
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
