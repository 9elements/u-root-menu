package submenu

import (
	"strings"

	"github.com/alimsk/list"

	"github.com/9elements/u-root-menu/pkg/menu"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Viewport *viewport.Model
	Help     help.Model

	Menu menu.BIOSMenu
	list list.Model

	filterPattern string

	keyMap keyMap
}

func Setup(menu menu.BIOSMenu, viewport *viewport.Model) *Model {
	// Set up the list
	listModel := list.NewSimpleAdapter(list.SimpleItemList{
		{
			Title:          "First Item",
			Desc:           "Some helper test here that guides you",
			Options:        []string{"Enabled", "Disabled"},
			SelectedOption: "Enabled",
			Disabled:       true,
		},
		{
			Title:          "Second Item",
			Desc:           "Some helper test here that guides you",
			Options:        []string{"Enabled", "Disabled"},
			SelectedOption: "Disabled",
			Disabled:       false,
		},
		{
			Title:          "Third Item",
			Desc:           "Some helper test here that guides you",
			Options:        []string{"Enabled", "Disabled"},
			SelectedOption: "Enabled",
			Disabled:       false,
		},
		{
			Title:          "Fourth Item",
			Desc:           "Some helper test here that guides you",
			Options:        []string{"Enabled", "Disabled"},
			SelectedOption: "Enabled",
			Disabled:       false,
		},
		{
			Title:          "Fifth Item",
			Desc:           "Some helper test here that guides you",
			Options:        []string{"Enabled", "Disabled"},
			SelectedOption: "Enabled",
			Disabled:       false,
		},
	})

	return &Model{
		Menu:     menu,
		list:     list.New(listModel),
		Viewport: viewport,
	}
}

func (m *Model) Init() tea.Cmd {
	m.list.Focus()
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit
			// case key.Matches(msg, m.keyMap.Enter):
			// 	adapter := m.list.Adapter.(*list.SimpleAdapter)
			// 	item := adapter.FilteredItemAt(m.list.ItemFocus())
			// 	if !item.Disabled {
			// 		return m, tea.Quit
			// 	}
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	menuDoc := strings.Builder{}

	ws := lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("39")).
		PaddingLeft(20).
		PaddingRight(20).
		PaddingTop(2).
		Align(lipgloss.Left).
		Border(lipgloss.RoundedBorder()).
		Width(m.Viewport.Width - 7)

	// Render the menu
	menuItems := m.list.View()
	menuDoc.WriteString(menuItems)

	docHeight := strings.Count(menuDoc.String(), "\n")
	requiredNewlinesForPadding := m.Viewport.Height - docHeight - 13

	menuDoc.WriteString(strings.Repeat("\n", max(0, requiredNewlinesForPadding)))

	return ws.Render(menuDoc.String())
}
