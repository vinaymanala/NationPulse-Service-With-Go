package utils

import (
	"context"

	"github.com/nationpulse-bff/internal/store"
)

type Configs struct {
	Db      *store.PgClient
	Cache   *store.Redis
	Context context.Context
}
