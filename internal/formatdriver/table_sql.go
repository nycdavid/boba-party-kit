package formatdriver

import "encoding/json"

type (
	TableSQL struct{}
)

func NewTableSQL(columns []string) *TableSQL {
	return &TableSQL{}
}

func (t *TableSQL) Format(data []byte) ([][]string, []string, error) {
	var d []map[string]string

	if err := json.Unmarshal(data, &d); err != nil {
		return nil, nil, err
	}

	var cols []string
	rows := make([][]string, len(d))

	for i, r := range d {
		if i == 0 {
			for k, _ := range r {
				cols = append(cols, k)
			}
		}

		rows[i] = make([]string, len(r))
		j := 0
		for _, v := range r {
			rows[i][j] = v
			j++
		}
	}

	return rows, cols, nil
}
