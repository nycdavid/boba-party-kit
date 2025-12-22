package table

import (
	btable "github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

type (
	View struct {
		baseStyle lipgloss.Style
		style     btable.Styles
	}
)

func NewView() *View {
	baseStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))

	s := btable.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	return &View{style: s, baseStyle: baseStyle}
}

func (t *View) Render(headers []string, rows [][]string) string {
	columns := headersToColumns(headers)
	tableRows := rowsToTableRows(rows)

	tbl := btable.New(
		btable.WithColumns(columns),
		btable.WithRows(tableRows),
		btable.WithFocused(true),
		btable.WithHeight(40),
	)
	tbl.SetStyles(t.style)

	return t.baseStyle.Render(tbl.View() + "\n")
}

func headersToColumns(headers []string) []btable.Column {
	columns := make([]btable.Column, len(headers))
	for i, header := range headers {
		columns[i] = btable.Column{Title: header, Width: 20}
	}

	return columns
}

func rowsToTableRows(rows [][]string) []btable.Row {
	tableRows := make([]btable.Row, len(rows))
	for i, row := range rows {
		tableRows[i] = row
	}

	return tableRows
}
