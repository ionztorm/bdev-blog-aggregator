package state

import (
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
)

type State struct {
	DB     *database.Queries
	Cfg    *config.Config
	DBConn *sql.DB
}
