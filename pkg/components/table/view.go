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
		width       int
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

	return &View{style: s, baseStyle: baseStyle, borderColor: ui.InactiveColor}
}

func (v *View) Render(headers []string, rows [][]string) string {
	columns := v.headersToColumns(headers)
	tableRows := rowsToTableRows(rows)

	tbl := btable.New(
		btable.WithColumns(columns),
		btable.WithRows(tableRows),
		btable.WithFocused(true),
		btable.WithHeight(40),
	)
	tbl.SetStyles(v.style)

	return v.baseStyle.BorderForeground(v.borderColor).Render(tbl.View() + "\n")
}

func (v *View) headersToColumns(headers []string) []btable.Column {
	columns := make([]btable.Column, len(headers))
	for i, header := range headers {
		columns[i] = btable.Column{Title: header, Width: (v.width - 10) / len(headers)}
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
