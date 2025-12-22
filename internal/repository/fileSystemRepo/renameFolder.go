package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) RenameFolder(
	ctx context.Context,
	userID uuid.UUID,
	folderID uuid.UUID,
	newName string,
) error {
	cmd, err := r.pool.Exec(
		ctx,
		queryRenameFolder,
		userID,
		folderID,
		newName,
	)
	if err != nil {
		return mapRepoError(err)
	}

	if cmd.RowsAffected() == 0 {
		// папка не найдена или не принадлежит пользователю
		return domain.ErrNotFound
	}

	return nil
}
