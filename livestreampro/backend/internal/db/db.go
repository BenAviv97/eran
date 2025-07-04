// /home/${USER}/livestreampro/backend/internal/db/db.go
package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Connect opens a PostgreSQL database using the provided connection string.
// The caller is responsible for closing the returned *sql.DB.
func Connect(connStr string) (*sql.DB, error) {
	return sql.Open("postgres", connStr)
}
