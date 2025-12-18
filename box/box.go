package box

import (
	"database/sql"

	"github.com/chains-lab/kafkakit/box/pgdb"
)

type Box struct {
	queries *pgdb.Queries
	db      *sql.DB
}

func New(db *sql.DB) Box {
	return Box{
		queries: pgdb.New(db),
		db:      db,
	}
}
