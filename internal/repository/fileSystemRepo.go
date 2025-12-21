package repository

import (
	"context"
	"file-service/internal/domain"

	"github.com/google/uuid"
)

// FileSystemRepo defines persistence operations for a user-scoped
// hierarchical file system (folders + files stored in S3, metadata in DB).
//
// All methods are strictly scoped by userID.
// Passing a valid context.Context is required; methods must respect
// cancellation and deadlines.
type FileSystemRepo interface {

	// CreateFolder creates a new folder for the given user.
	//
	// If parentFolderID is nil, the folder is created in the root.
	// The folder name must be unique within the same parent folder
	// for the given user.
	//
	// Returns the created Folder or an error if:
	//   - a folder with the same name already exists in the target parent
	//   - the parent folder does not exist
	//   - the context is cancelled or times out
	CreateFolder(
		ctx context.Context,
		userID uuid.UUID,
		parentFolderID *uuid.UUID,
		folderName string,
	) (domain.Folder, error)

	// GetFolder returns a folder by its ID for the given user.
	//
	// Returns an error if the folder does not exist or does not
	// belong to the specified user.
	GetFolder(
		ctx context.Context,
		userID uuid.UUID,
		folderID uuid.UUID,
	) (domain.Folder, error)

	// ListFolders returns all direct child folders of the given parent folder.
	//
	// If parentID is nil, folders from the root level are returned.
	// The result is scoped to the specified user.
	ListFolders(
		ctx context.Context,
		userID uuid.UUID,
		parentID *uuid.UUID,
	) ([]domain.Folder, error)

	// RenameFolder renames an existing folder.
	//
	// The new name must be unique within the same parent folder
	// for the given user.
	//
	// Returns an error if:
	//   - the folder does not exist
	//   - a name conflict occurs
	RenameFolder(
		ctx context.Context,
		userID uuid.UUID,
		folderID uuid.UUID,
		newName string,
	) error

	// MoveFolder moves a folder to a new parent folder.
	//
	// If newParentID is nil, the folder is moved to the root.
	//
	// Implementations should prevent invalid moves, such as:
	//   - moving a folder into itself
	//   - moving a folder into one of its descendants
	//
	// Returns an error if the target parent does not exist or
	// the move violates folder hierarchy constraints.
	MoveFolder(
		ctx context.Context,
		userID uuid.UUID,
		folderID uuid.UUID,
		newParentID *uuid.UUID,
	) error

	// DeleteFolder permanently deletes a folder and all its subfolders
	// for the given user.
	//
	// Associated file metadata is also deleted according to the database
	// foreign key constraints (ON DELETE CASCADE).
	DeleteFolder(
		ctx context.Context,
		userID uuid.UUID,
		folderID uuid.UUID,
	) error

	// FolderExists checks whether a folder exists for the given user.
	//
	// Returns true if the folder exists, false otherwise.
	// This method does not return the folder itself.
	FolderExists(
		ctx context.Context,
		userID uuid.UUID,
		folderID uuid.UUID,
	) (bool, error)

	// CreateFileMeta creates a metadata record for a file that has started uploading.
	//
	// The file is created with status "uploading".
	// The actual file contents are expected to be uploaded to S3 separately.
	//
	// If folderID is nil, the file is placed in the root folder.
	//
	// Returns the created File metadata or an error if:
	//   - a file with the same name already exists in the target folder
	//   - the provided S3 key is not unique for the user
	CreateFileMeta(
		ctx context.Context,
		userID uuid.UUID,
		folderID *uuid.UUID,
		fileName string,
		s3Bucket string,
		s3Key string,
		mimeType string,
		size int64,
	) (domain.File, error)

	// GetFileMeta returns file metadata by file ID for the given user.
	//
	// Returns an error if the file does not exist or does not
	// belong to the specified user.
	GetFileMeta(
		ctx context.Context,
		userID uuid.UUID,
		fileID uuid.UUID,
	) (domain.File, error)

	// GetFileMetaByKey returns file metadata by its S3 object key.
	//
	// This method is useful for handling background jobs,
	// callbacks, or S3-related workflows.
	GetFileMetaByKey(
		ctx context.Context,
		userID uuid.UUID,
		s3Key string,
	) (domain.File, error)

	// ListFilesMeta returns all files located in the given folder.
	//
	// If folderID is nil, files from the root folder are returned.
	// The result is scoped to the specified user.
	ListFilesMeta(
		ctx context.Context,
		userID uuid.UUID,
		folderID *uuid.UUID,
	) ([]domain.File, error)

	// RenameFile renames an existing file.
	//
	// The new name must be unique within the same folder
	// for the given user.
	RenameFile(
		ctx context.Context,
		userID uuid.UUID,
		fileID uuid.UUID,
		newName string,
	) error

	// MoveFile moves a file to a different folder.
	//
	// If newFolderID is nil, the file is moved to the root folder.
	MoveFile(
		ctx context.Context,
		userID uuid.UUID,
		fileID uuid.UUID,
		newFolderID *uuid.UUID,
	) error

	// DeleteFileMeta permanently deletes file metadata for the given user.
	//
	// The actual file contents in S3 must be deleted separately
	// by the application or background worker.
	DeleteFileMeta(
		ctx context.Context,
		userID uuid.UUID,
		fileID uuid.UUID,
	) error

	// CommitFileUpload marks a file upload as successfully completed.
	//
	// This method transitions the file status from "uploading" to "uploaded".
	// It should fail if the file is already uploaded or does not exist.
	CommitFileUpload(
		ctx context.Context,
		userID uuid.UUID,
		fileID uuid.UUID,
	) error
}
