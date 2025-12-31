package formatdriver

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/itchyny/gojq"
	"github.com/nycdavid/boba-party-kit/internal/config"
)

type (
	TableJSON struct {
		rowRoot string
		columns []config.JSONColumn
	}
)

func NewTableJSON(rowRoot string, c []config.JSONColumn) *TableJSON {
	return &TableJSON{rowRoot: rowRoot, columns: c}
}

func (j *TableJSON) Format(data []byte) ([][]string, []string, error) {
	var jsonData map[string]any
	var rows [][]string

	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, nil, err
	}

	colFragments := make([]string, len(j.columns))
	for i, col := range j.columns {
		colFragments[i] = fmt.Sprintf("%s: %s", col.Name, col.Path)
	}
	joinCols := "{" + strings.Join(colFragments, ", ") + "}"

	q, err := gojq.Parse(j.rowRoot + " | " + joinCols)
	if err != nil {
		return nil, nil, err
	}

	var lines []map[string]any
	iter := q.Run(jsonData)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}

		lines = append(lines, v.(map[string]any))
	}

	for _, line := range lines {
		rowCells := make([]string, len(j.columns))
		for i, col := range j.columns {
			switch val := line[col.Name].(type) {
			case string:
				rowCells[i] = val
			case float64:
				rowCells[i] = strconv.FormatFloat(val, 'f', 2, 64)
			case nil:
				rowCells[i] = ""
			}
		}

		rows = append(rows, rowCells)
	}

	columns := make([]string, len(j.columns))
	for i, col := range j.columns {
		columns[i] = col.Name
	}

	return rows, columns, nil
}
