package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

func (r *repo) GetFolder(
	ctx context.Context,
	userID uuid.UUID,
	folderID uuid.UUID,
) (domain.Folder, error) {
	var folder domain.Folder

	err := r.pool.QueryRow(
		ctx,
		queryGetFolder,
		userID,
		folderID,
	).Scan(
		&folder.ID,
		&folder.UserID,
		&folder.ParentID,
		&folder.Name,
		&folder.CreatedAt,
	)

	if err != nil {
		return domain.Folder{}, mapRepoError(err)
	}

	return folder, nil
}
