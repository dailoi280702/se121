package sqlutils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

func CountRecords(db *sql.DB, table string) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)

	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count records: %v", err)
	}

	return count, nil
}

func Exists(db *sql.DB, query string, args ...any) (bool, error) {
	exists := false
	if err := db.QueryRow(query, args...).Scan(&exists); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return false, nil
		default:
			log.Println(args...)
			return false, err
		}
	}
	return exists, nil
}

// arguments is slice of key-value, for example: k1 v1 k2 v2 k3 v3 k4 v4
func ScanRecordById(db *sql.DB, table string, id any, args ...any) error {
	if len(args)%2 != 0 {
		return errors.New("number of arguments must be even")
	}

	var keys []string
	var vals []any
	for i := 0; i < len(args); i += 1 {
		if i%2 == 0 {
			keys = append(keys, args[i].(string))
			continue
		}
		vals = append(vals, args[i])
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", strings.Join(keys, ", "), table)
	if err := db.QueryRow(query, id).Scan(vals...); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return fmt.Errorf("failed to insert record: %v", err)
		default:
			return err
		}
	}
	return nil
}

func UpdateRecord(db *sql.DB, table string, record map[string]interface{}, id int) error {
	query := fmt.Sprintf("UPDATE %s SET", table)

	args := make([]interface{}, 0)
	i := 1
	for key, value := range record {
		if value == nil {
			continue
		}
		query += fmt.Sprintf(" %s = $%d,", key, i)
		args = append(args, value)
		i++
	}
	query = query[:len(query)-1] + fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	_, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update record: %v", err)
	}

	return nil
}

func UpdateRecordWithTransaction(tx *sql.Tx, table string, record map[string]interface{}, id int) error {
	query := fmt.Sprintf("UPDATE %s SET", table)

	args := make([]interface{}, 0)
	i := 1
	for key, value := range record {
		if value == nil {
			continue
		}
		query += fmt.Sprintf(" %s = $%d,", key, i)
		args = append(args, value)
		i++
	}
	query = query[:len(query)-1] + fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	_, err := tx.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update record: %v", err)
	}

	return nil
}

func IdExists(db *sql.DB, table, id any) (bool, error) {
	query := `
        SELECT exists( SELECT 1 FROM %s WHERE id = $1)
        `
	query = fmt.Sprintf(query, table)
	return Exists(db, query, id)
}
