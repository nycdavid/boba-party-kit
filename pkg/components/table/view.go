package table

import (
	btable "github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/nycdavid/boba-party-kit/pkg/components/ui"
)

type (
	View struct {
		baseStyle   lipgloss.Style
		borderColor lipgloss.Color
		style       btable.Styles
		tbl         *btable.Model
	}
)

func NewView() *View {
	baseStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder())

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
	tbl := btable.New(
		btable.WithFocused(true),
		btable.WithWidth(200),
		btable.WithHeight(40),
	)
	tbl.SetStyles(s)

	return &View{style: s, baseStyle: baseStyle, borderColor: ui.InactiveColor, tbl: &tbl}
}

func (v *View) Render() string {
	return v.baseStyle.BorderForeground(v.borderColor).Render(v.tbl.View() + "\n")
}

func (v *View) headersToColumns(headers []string) []btable.Column {
	columns := make([]btable.Column, len(headers))
	for i, header := range headers {
		columns[i] = btable.Column{Title: header, Width: v.tbl.Width() / len(headers)}
	}

	return columns
}

func (v *View) resetColumnsWithNewWidth(tblWidth int) {
	cols := v.tbl.Columns()
	for _, col := range cols {
		col.Width = tblWidth / len(cols)
	}

	v.tbl.SetColumns(cols)
}

func (v *View) setColumns(columns []string) {
	v.tbl.SetColumns(v.headersToColumns(columns))
}

func (v *View) setRows(rows [][]string) {
	v.tbl.SetRows(rowsToTableRows(rows))
}

func rowsToTableRows(rows [][]string) []btable.Row {
	tableRows := make([]btable.Row, len(rows))
	for i, row := range rows {
		tableRows[i] = row
	}

	return tableRows
}
