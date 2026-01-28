package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // ← ЭТОТ ИМПОРТ ОБЯЗАТЕЛЕН!
)

/*
# 2. Проверить таблицу
docker exec -it postgres-storage psql -U postgres -d storage -c "\dt"

# 3. Проверить структуру таблицы
docker exec -it postgres-storage psql -U postgres -d storage -c "\d url"
*/
type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	stmt, err := db.Prepare(
		`
		CREATE TABLE IF NOT EXISTS url(
		id SERIAL PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL
		)
		`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt2, err := db.Prepare(`
    CREATE INDEX IF NOT EXISTS idx_alias ON url(alias)
`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt2.Close()

	_, err = stmt2.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil

}
