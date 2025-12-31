package datadriver

import (
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

type (
	SQLite struct {
		fpath string
		query string
	}
)

func NewSQLite(fpath, query string) *SQLite {
	return &SQLite{fpath: fpath, query: query}
}

func (s *SQLite) Fetch() ([]byte, error) {
	db, err := sql.Open("sqlite3", s.fpath)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(s.query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var rs []map[string]string
	for rows.Next() {
		values := make([]any, len(cols))
		valuePtrs := make([]any, len(cols))

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]string, len(cols))
		for i, col := range cols {
			val := values[i]

			if b, ok := val.([]byte); ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = val.(string)
			}
		}

		rs = append(rs, rowMap)
	}

	return json.Marshal(rs)
}
