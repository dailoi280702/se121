package service

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

var UnplementedError = errors.New("method unimplemented yet")

type ValidationErrors struct {
	errors map[string]string
}

func (e *ValidationErrors) Error() string {
	return fmt.Sprintf("%v", e.errors)
}

type Service struct {
	DB *sql.DB
}

func GetObjectFromDB[T interface{}](db *sql.DB, query string, args ...any) (*T, error) {
	t := new(T)
	s := reflect.ValueOf(t).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}

	if err := db.QueryRow(query, args...).Scan(columns...); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, nil
		default:
			return nil, err
		}
	}

	return t, nil
}

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
