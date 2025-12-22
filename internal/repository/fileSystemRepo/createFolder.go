package fileSystemRepo

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

// CreateFolder creates a new folder for the given user.
// parentFolderID == nil means root.
func (r *repo) CreateFolder(
	ctx context.Context,
	userID uuid.UUID,
	parentFolderID *uuid.UUID,
	folderName string,
) (domain.Folder, error) {
	var folder domain.Folder

	err := r.pool.QueryRow(
		ctx,
		queryCreateFolder,
		userID,
		parentFolderID,
		folderName,
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
