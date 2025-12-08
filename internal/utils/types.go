package utils

import (
	"context"
	"database/sql"

	"github.com/nationpulse-bff/internal/store"
)

type Configs struct {
	Db      *sql.DB
	Cache   *store.Redis
	Context context.Context
}
