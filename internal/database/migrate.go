package database

import (
	"database/sql"
	"embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed sql/schema/*.sql
var migrationsFS embed.FS

func RunMigrations(db *sql.DB, direction string) error {
	direction = strings.ToLower(direction)
	if direction != "up" && direction != "down" {
		return fmt.Errorf("invalid migration direction %q: must be 'up' or 'down'", direction)
	}

	entries, err := migrationsFS.ReadDir("sql/schema")
	if err != nil {
		return fmt.Errorf("reading migrations directory: %w", err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, entry.Name())
		}
	}

	sort.Strings(files)
	if direction == "down" {
		for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
			files[i], files[j] = files[j], files[i]
		}
	}

	for _, file := range files {
		path := "sql/schema/" + file
		data, err := migrationsFS.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading migration file %s: %w", file, err)
		}

		sqlSection, err := extractSQLSection(string(data), direction)
		if err != nil {
			return fmt.Errorf("extracting %s section from %s: %w", direction, file, err)
		}

		if sqlSection == "" {
			fmt.Printf("skipping migration %s: no %s section found\n", file, direction)
			continue
		}

		fmt.Printf("Running migration %s: %s\n", direction, file)
		if _, err := db.Exec(sqlSection); err != nil {
			return fmt.Errorf("failed migration %s (%s): %w", file, direction, err)
		}
	}

	fmt.Printf("All %s migrations applied successfully\n", direction)
	return nil
}

func extractSQLSection(sqlContent, direction string) (string, error) {
	marker := fmt.Sprintf("-- %s", direction)
	lines := strings.Split(sqlContent, "\n")

	var (
		inSection bool
		section   []string
	)

	for _, line := range lines {
		lineTrim := strings.TrimSpace(line)

		if strings.HasPrefix(lineTrim, "-- up") || strings.HasPrefix(lineTrim, "-- down") {
			if lineTrim == marker {
				inSection = true
			} else if inSection {
				// End of current section when a new section starts
				break
			}
			continue
		}

		if inSection {
			section = append(section, line)
		}
	}

	return strings.TrimSpace(strings.Join(section, "\n")), nil
}
