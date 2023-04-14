package db_store

import (
	"database/sql"
	"errors"
)

func ExistInDB(db *sql.DB, query string, args ...any) (bool, error) {
	existed := false

	if err := db.QueryRow(query, args...).Scan(&existed); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return false, nil
		default:
			return false, err
		}
	}

	return existed, nil
}
