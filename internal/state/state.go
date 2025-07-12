package state

import (
	"database/sql"

	"github.com/ionztorm/gator/internal/config"
	"github.com/ionztorm/gator/internal/database"
)

type State struct {
	DB     *database.Queries
	Cfg    *config.Config
	DBConn *sql.DB
}
