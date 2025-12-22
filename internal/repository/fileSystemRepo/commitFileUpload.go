package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) CommitFileUpload(
	ctx context.Context,
	userID uuid.UUID,
	fileID uuid.UUID,
) error {
	cmd, err := r.pool.Exec(
		ctx,
		queryCommitFileUpload,
		userID,
		fileID,
	)
	if err != nil {
		return mapRepoError(err)
	}

	if cmd.RowsAffected() == 0 {
		// Возможные причины:
		// - файл не существует
		// - файл принадлежит другому пользователю
		// - файл уже в статусе 'uploaded'
		return domain.ErrInvalidState
	}

	return nil
}
