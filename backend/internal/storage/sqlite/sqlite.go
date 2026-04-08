package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/sapienfrom2000s/trident/backend/internal/core"
)

type Sqlite struct {
	DB *sql.DB
}

func New(db *sql.DB) Sqlite {
	return Sqlite{}
}

func (s Sqlite) StoreEvent(n core.NormalizedEvent) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func (s Sqlite) StoreJob(j core.Job) (bool, error) {
	return false, fmt.Errorf("not implemented")
}
