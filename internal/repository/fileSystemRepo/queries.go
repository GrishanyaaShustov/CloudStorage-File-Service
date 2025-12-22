package fileSystemRepo

const (
	queryCreateFolder = `
	INSERT INTO folders (
		user_id,
		parent_id,
		name
	)
	VALUES (
		$1,     -- user_id
		$2,     -- parent_id (NULL = корень)
		$3      -- folder name
	)
	RETURNING
		id,
		user_id,
		parent_id,
		name,
		created_at;
	`

	queryGetFolder = `
	SELECT
		id,
		user_id,
		parent_id,
		name,
		created_at
	FROM folders
	WHERE
		user_id = $1
		AND id = $2;
	`

	queryListFolders = `
	SELECT
		id,
		user_id,
		parent_id,
		name,
		created_at
	FROM folders
	WHERE
		user_id = $1
		AND (
			($2 IS NULL AND parent_id IS NULL)
			OR
			(parent_id = $2)
		)
	ORDER BY name ASC;
	`

	queryRenameFolder = `
	UPDATE folders
	SET
		name = $3
	WHERE
		user_id = $1
		AND id = $2;
	`

	queryMoveFolder = `
	UPDATE folders
	SET
		parent_id = $3
	WHERE
		user_id = $1
		AND id = $2
		AND (
			$3 IS NULL OR $3 <> $2
		);
	`

	queryDeleteFolder = `
	DELETE FROM folders
	WHERE
		user_id = $1
		AND id = $2;
	`

	queryFolderExists = `
	SELECT EXISTS (
		SELECT 1
		FROM folders
		WHERE
			user_id = $1
			AND id = $2
	);
	`

	queryCreateFileMeta = `
	INSERT INTO files (
		user_id,
		folder_id,
		name,
		s3_bucket,
		s3_key,
		mime_type,
		size_bytes,
		status
	)
	VALUES (
		$1,     -- user_id
		$2,     -- folder_id (NULL = корень)
		$3,     -- file name
		$4,     -- s3 bucket
		$5,     -- s3 key
		$6,     -- mime type
		$7,     -- size bytes
		'uploading'
	)
	RETURNING
		id,
		user_id,
		folder_id,
		name,
		s3_bucket,
		s3_key,
		mime_type,
		size_bytes,
		status,
		created_at,
		updated_at;
	`

	queryGetFileMeta = `
	SELECT
		id,
		user_id,
		folder_id,
		name,
		s3_bucket,
		s3_key,
		mime_type,
		size_bytes,
		status,
		created_at,
		updated_at
	FROM files
	WHERE
		user_id = $1
		AND id = $2;
	`

	queryGetFileMetaByKey = `
	SELECT
		id,
		user_id,
		folder_id,
		name,
		s3_bucket,
		s3_key,
		mime_type,
		size_bytes,
		status,
		created_at,
		updated_at
	FROM files
	WHERE
		user_id = $1
		AND s3_key = $2;
	`

	queryListFilesMeta = `
	SELECT
		id,
		user_id,
		folder_id,
		name,
		s3_bucket,
		s3_key,
		mime_type,
		size_bytes,
		status,
		created_at,
		updated_at
	FROM files
	WHERE
		user_id = $1
		AND (
			($2 IS NULL AND folder_id IS NULL)
			OR
			(folder_id = $2)
		)
	ORDER BY name ASC;
	`

	queryRenameFile = `
	UPDATE files
	SET
		name = $3,
		updated_at = now()
	WHERE
		user_id = $1
		AND id = $2;
	`

	queryMoveFile = `
	UPDATE files
	SET
		folder_id = $3,
		updated_at = now()
	WHERE
		user_id = $1
		AND id = $2;
	`

	queryDeleteFileMeta = `
	DELETE FROM files
	WHERE
		user_id = $1
		AND id = $2;
	`

	queryCommitFileUpload = `
	UPDATE files
	SET
		status = 'uploaded',
		updated_at = now()
	WHERE
		user_id = $1
		AND id = $2
		AND status = 'uploading';
	`
)
