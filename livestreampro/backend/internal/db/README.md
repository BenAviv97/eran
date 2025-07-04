Package db contains a minimal PostgreSQL client using database/sql.
It centralizes connection logic so other packages can obtain *sql.DB
instances without managing driver specifics.
