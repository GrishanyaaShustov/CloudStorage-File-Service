package domain

import (
	"time"

	"github.com/google/uuid"
)

type Folder struct {
	ID        uuid.UUID  // id папки (uuidv7)
	UserID    uuid.UUID  // владелец
	ParentID  *uuid.UUID // nil = корень
	Name      string     // имя папки
	CreatedAt time.Time  // дата создания
}
