CREATE TABLE IF NOT EXISTS files (
                                     user_id         UUID NOT NULL,
                                     id              UUID NOT NULL DEFAULT uuidv7(),

                                     folder_id       UUID NULL,
                                     name            TEXT NOT NULL,

                                     s3_bucket       TEXT NOT NULL,
                                     s3_key          TEXT NOT NULL,

                                     content_type    TEXT NULL,
                                     size_bytes      BIGINT NOT NULL CHECK (size_bytes >= 0),
                                     sha256          TEXT NULL,
                                     etag            TEXT NULL,

                                     created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
                                     updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),

                                     PRIMARY KEY (user_id, id),

                                     FOREIGN KEY (user_id, folder_id)
                                         REFERENCES folders(user_id, id)
                                         ON DELETE SET NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS ux_files_user_folder_name
    ON files (user_id, folder_id, name);

CREATE UNIQUE INDEX IF NOT EXISTS ux_files_user_s3key
    ON files (user_id, s3_key);

CREATE INDEX IF NOT EXISTS idx_files_user_folder
    ON files (user_id, folder_id);

CREATE INDEX IF NOT EXISTS idx_files_user_created_at
    ON files (user_id, created_at);