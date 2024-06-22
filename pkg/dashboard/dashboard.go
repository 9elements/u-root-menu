package dashboard

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const banner = `
MMMMMMMMMMMMMMMWo                                    oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo                                    oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo                                    oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo                                    oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo         :kd,         .;ox:.        oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo         ,xKXk;.   .;d0XOc.         oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo           .l0XOodk0XOo,.           oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo             ,OWMMMNd.              oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo           .l0N0dxOKXk;.            oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo         'dKXk:.   .:kKOc.          oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo         c0d'         'col,.        oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo         ..               .         oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo        .:oolllllllllllool;         oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo        .cdxxxxxxxxxxdddol,         oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo                                    oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo                                    oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWo                                    oWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWk;'','','''''''''''''''''''''''',,,';kWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMWWWWNNWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
`
const welcomeString = "Welcome to u-root-menu"

type Model struct {
	Viewport *viewport.Model
	Help     help.Model

	keyMap keyMap
}

func Setup() *Model {
	return &Model{}
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
		}
	}

	return m, cmd
}

func (m *Model) View() string {
	dashboardDoc := strings.Builder{}

	ws := lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("39")).
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		Width(m.Viewport.Width - 7)

	dashboardDoc.WriteString(lipgloss.JoinVertical(lipgloss.Center, banner, welcomeString))

	docHeight := strings.Count(dashboardDoc.String(), "\n")
	requiredNewlinesForPadding := m.Viewport.Height - docHeight - 13

	dashboardDoc.WriteString(strings.Repeat("\n", max(0, requiredNewlinesForPadding)))

	return ws.Render(dashboardDoc.String())
}

func (m *Model) Init() tea.Cmd {
	return nil
}
