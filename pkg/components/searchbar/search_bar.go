package searchbar

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	SearchBar struct {
		textInput textinput.Model
	}

	Mod func(*SearchBar)
)

func New(opts ...Mod) *SearchBar {
	s := &SearchBar{}

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
	return s, nil
}

func (s *SearchBar) View() string {
	return s.textInput.View()
}
