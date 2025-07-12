package command

import (
	"fmt"
	"log"

	"github.com/ionztorm/gator/internal/database"
	"github.com/ionztorm/gator/internal/state"
)

func init() {
	registerGlobal("migrate", HandleMigrate)
}

func HandleMigrate(s *state.State, cmd Command) error {
	direction := "up"

	if len(cmd.Args) > 0 {
		arg := cmd.Args[0]
		if arg != "up" && arg != "down" {
			return fmt.Errorf("invalid migration direction: must be 'up' or 'down'")
		}
		direction = arg
	}

	if s.DBConn == nil {
		return fmt.Errorf("DB connection is nil")
	}

	err := database.RunMigrations(s.DBConn, direction)
	if err != nil {
		log.Printf("Migration error: %v", err)
		return fmt.Errorf("migration failed: %w", err)
	}

	fmt.Println("✅ Migrations completed successfully")
	return nil
}
