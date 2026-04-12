package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/sapienfrom2000s/trident/backend/internal/core/models"
)

type Sqlite struct {
	DB *sql.DB
}

func New(db *sql.DB) Sqlite {
	return Sqlite{}
}

func (s Sqlite) StoreEvent(n models.Event) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func (s Sqlite) StoreJob(j models.Job) (bool, error) {
	return false, fmt.Errorf("not implemented")
}
