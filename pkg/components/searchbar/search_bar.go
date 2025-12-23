package searchbar

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nycdavid/boba-party-kit/pkg/components/ui"
)

type (
	SearchBar struct {
		baseStyle   lipgloss.Style
		borderColor lipgloss.Color
		textInput   textinput.Model
	}

	Mod func(*SearchBar)
)

func New(opts ...Mod) *SearchBar {
	s := &SearchBar{
		baseStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()),
		borderColor: ui.ActiveColor,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.textInput = textinput.New()
	s.textInput.Placeholder = "Start typing to search..."
	s.textInput.Focus()
	s.textInput.CharLimit = 156
	s.textInput.Width = 80

	return s
}

func (s *SearchBar) Init() tea.Cmd {
	return textinput.Blink
}

func (s *SearchBar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ui.LoseFocusMsg:
		s.borderColor = ui.InactiveColor
	case ui.TakeFocusMsg:
		s.borderColor = ui.ActiveColor
	case tea.KeyMsg:
		var cmd tea.Cmd
		s.textInput, cmd = s.textInput.Update(msg)
		return s, cmd
	}

	return s, nil
}

func (s *SearchBar) View() string {
	return s.baseStyle.BorderForeground(s.borderColor).Render(s.textInput.View())
}
