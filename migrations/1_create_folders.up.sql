CREATE TABLE IF NOT EXISTS folders (
                                       user_id     UUID NOT NULL,
                                       id          UUID NOT NULL DEFAULT uuidv7(),
                                       parent_id   UUID NULL,

                                       name        TEXT NOT NULL,
                                       created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

                                       PRIMARY KEY (user_id, id),

                                       FOREIGN KEY (user_id, parent_id)
                                           REFERENCES folders(user_id, id)
                                           ON DELETE CASCADE
);


CREATE UNIQUE INDEX IF NOT EXISTS ux_folders_user_parent_name
    ON folders (user_id, parent_id, name);

CREATE INDEX IF NOT EXISTS idx_folders_user_parent
    ON folders (user_id, parent_id);