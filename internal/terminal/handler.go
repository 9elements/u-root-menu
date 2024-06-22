package terminal

import (
	"fmt"
	"strings"

	ts "github.com/9elements/u-root-menu/internal/style"
	"github.com/9elements/u-root-menu/pkg/boot"
	"github.com/9elements/u-root-menu/pkg/dashboard"
	"github.com/9elements/u-root-menu/pkg/menu"
	"github.com/9elements/u-root-menu/pkg/submenu"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	version string

	TabsWithColors []string
	currentTab     *int

	terminalSizeReady bool

	dashboardInfo tea.Model
	bootSelection tea.Model

	categories map[string]tea.Model

	keyMap keyMap

	viewport viewport.Model
}

func Setup(version string, biosMenu menu.BIOSMenu) tea.Model {
	currentTab := new(int)

	// List of categories here - Dashboard might always be there, rest is optional
	tabsWithColor := []string{"Dashboard", "Boot Menu"}

	submenues := biosMenu.Categories()
	categories := make(map[string]tea.Model)

	dashboardInfo := dashboard.Setup()
	bootSelection := boot.Setup()

	m := model{
		version:        version,
		TabsWithColors: tabsWithColor,
		currentTab:     currentTab,

		dashboardInfo: dashboardInfo,
		bootSelection: bootSelection,

		keyMap: keys,
	}

	// Build up the categories dynamically
	for name := range submenues {
		tabsWithColor = append(tabsWithColor, name)

		if name == "Boot Menu" {
			categories[name] = submenu.Setup(submenues[name], &m.viewport)
		}
	}

	m.categories = categories
	m.TabsWithColors = tabsWithColor

	dashboardInfo.Viewport = &m.viewport
	bootSelection.Viewport = &m.viewport

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
	var renderedTitles string
	for _, t := range titles {
		renderedTitles += t
	}
	line := strings.Repeat("─", max(0, m.viewport.Width-59))
	titles = append(titles, line)

	titles = append(titles, ts.TitleStyleDisabled.Copy().Render(m.version))

	return lipgloss.JoinHorizontal(lipgloss.Center, titles...)
}

func (m *model) Init() tea.Cmd {
	// Initialize all the categories
	for _, category := range m.categories {
		category.Init()
	}

	return tea.Batch(
		tea.EnterAltScreen,
		tea.SetWindowTitle("Firmware Menu"))
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.syncTerminal(msg)

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// If key matches tab switch
		case key.Matches(msg, m.keyMap.Tab):
			*m.currentTab = (*m.currentTab + 1) % len(m.TabsWithColors)
			cmds = append(cmds, m.handleTabContent(cmd, msg))
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit
		default:
			cmds = append(cmds, m.handleTabContent(cmd, msg))
		}
	}

	return m, tea.Batch(cmds...)
}

const (
	minTerminalWidth  = 0
	minTerminalHeight = 0
)

func (m *model) View() string {
	if !m.terminalSizeReady {
		return "Setting up..."
	}

	if m.viewport.Width < minTerminalWidth || m.viewport.Height < minTerminalHeight {
		return fmt.Sprintf("Terminal window is too small. Please resize to at least %dx%d.", minTerminalWidth, minTerminalHeight)
	}

	var mainDocument strings.Builder
	// var operationDocument string
	// var helpDocument string

	var renderedTabs []string

	for i, t := range m.TabsWithColors {
		var style lipgloss.Style
		if i == *m.currentTab {
			style = ts.TitleStyleActive
		} else {
			style = ts.TitleStyleInactive
		}

		renderedTabs = append(renderedTabs, style.Render(t))
	}

	mainDocument.WriteString(m.headerView(renderedTabs...) + "\n")

	switch *m.currentTab {
	case 0:
		mainDocument.WriteString(m.dashboardInfo.View())
	case 1:
		mainDocument.WriteString(m.bootSelection.View())
	default:
		if m.categories[m.TabsWithColors[*m.currentTab]] == nil {
			mainDocument.WriteString("Not implemented yet")
			break
		}

		// Check if the category is present in the menu
		_, ok := m.categories[m.TabsWithColors[*m.currentTab]]
		if !ok {
			mainDocument.WriteString("Category is not present.")
			break
		}

		// Render the category
		mainDocument.WriteString(m.categories[m.TabsWithColors[*m.currentTab]].View())
	}

	mainDocumentContent := ts.DocStyle.Render(mainDocument.String())

	return mainDocumentContent
}

func (m *model) handleTabContent(cmd tea.Cmd, msg tea.Msg) tea.Cmd {
	switch m.TabsWithColors[*m.currentTab] {
	case "Dashboard":
		m.dashboardInfo, cmd = m.dashboardInfo.Update(msg)
	case "Boot Menu":
		m.bootSelection.Init()
		m.bootSelection, cmd = m.bootSelection.Update(msg)
	default:
		if m.categories[m.TabsWithColors[*m.currentTab]] == nil {
			return cmd
		}

		// Check if the category is present in the menu
		_, ok := m.categories[m.TabsWithColors[*m.currentTab]]
		if !ok {
			return cmd
		}

		// Update the Model
		m.categories[m.TabsWithColors[*m.currentTab]], cmd = m.categories[m.TabsWithColors[*m.currentTab]].Update(msg)
	}
	return cmd
}
