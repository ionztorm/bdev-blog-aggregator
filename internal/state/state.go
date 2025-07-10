package state

import (
	"gator/internal/config"
	"gator/internal/database"
)

type State struct {
	DB  *database.Queries
	Cfg *config.Config
}
