package terminal

import (
	"strings"

	ts "github.com/9elements/u-root-menu/internal/style"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	TabsWithColors    []string
	currentTab        *int
	isTabActive       bool
	terminalSizeReady bool

	keyMap keyMap

	viewport viewport.Model
}

func Setup() tea.Model {
	currentTab := new(int)

	tabsWithColor := []string{"Dashboard", "Firmware"}

	m := model{
		TabsWithColors: tabsWithColor,
		currentTab:     currentTab,
		keyMap:         keys,
	}

	return &m
}

func (m *model) syncTerminal(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())

		if !m.terminalSizeReady {
			m.viewport = viewport.New(msg.Width, msg.Height)
			m.viewport.YPosition = headerHeight + 1
			m.terminalSizeReady = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height
		}
	}
}

func (m *model) headerView(titles ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Center, titles...)
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tea.SetWindowTitle("Firmware Menu"))
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.syncTerminal(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// If key matches tab switch
		case key.Matches(msg, m.keyMap.Tab):
			*m.currentTab = (*m.currentTab + 1) % len(m.TabsWithColors)
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *model) View() string {
	if !m.terminalSizeReady {
		return "Setting up..."
	}

	var mainDocument strings.Builder
	// var operationDocument string
	// var helpDocument string

	var renderedTabs []string

	for i, t := range m.TabsWithColors {
		var style lipgloss.Style
		if i == *m.currentTab {
			style = ts.TitleStyleActive.Copy()
		} else {
			style = ts.TitleStyleInactive.Copy()
		}

		renderedTabs = append(renderedTabs, style.Render(t))
	}

	mainDocument.WriteString(m.headerView(renderedTabs...))

	mainDocumentContent := ts.DocStyle.Render(mainDocument.String())

	return mainDocumentContent
}

func (m *model) handelTabContent(cmd tea.Cmd, msg tea.Msg) tea.Cmd {
	return cmd
}
