package style

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	DocStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	WindowStyleCyan   = lipgloss.NewStyle().BorderForeground(lipgloss.Color("39"))
	WindowStyleOrange = lipgloss.NewStyle().BorderForeground(lipgloss.Color("#ffaf00")).Border(lipgloss.RoundedBorder())
	WindowStyleRed    = lipgloss.NewStyle().BorderForeground(lipgloss.Color("9")).Border(lipgloss.RoundedBorder())
	WindowStyleGreen  = lipgloss.NewStyle().BorderForeground(lipgloss.Color("10")).Border(lipgloss.RoundedBorder())
	WindowStyleGray   = lipgloss.NewStyle().BorderForeground(lipgloss.Color("240")).Border(lipgloss.NormalBorder())
	WindowStyleWhite  = lipgloss.NewStyle().BorderForeground(lipgloss.Color("255")).Border(lipgloss.NormalBorder())
	WindowStyleYellow = lipgloss.NewStyle().BorderForeground(lipgloss.Color("11")).Border(lipgloss.NormalBorder())
	WindowStylePink   = lipgloss.NewStyle().BorderForeground(lipgloss.Color("205")).Border(lipgloss.RoundedBorder())

	WindowStyleHelp           = WindowStyleGray.Margin(0, 0, 0, 1).Padding(0, 2, 0, 2)
	WindowStyleFooter         = lipgloss.NewStyle().Margin(0, 0, 0, 1).Padding(0, 2, 0, 2).AlignHorizontal(lipgloss.Center)
	WindowStyleError          = WindowStyleRed.Margin(0, 0, 0, 1).Padding(0, 2, 0, 2)
	WindowStyleProgress       = WindowStyleOrange.Margin(0, 0, 0, 1).Padding(0, 2, 0, 2)
	WindowStyleSuccess        = WindowStyleGreen.Margin(0, 0, 0, 1).Padding(0, 2, 0, 2)
	WindowStyleDefault        = WindowStyleWhite.Margin(0, 0, 0, 1).Padding(0, 2, 0, 2)
	WindowStyleOptionSelector = WindowStylePink.Margin(0, 0, 0, 1).Padding(0, 2, 0, 2)
)

var (
	TitleStyleActive = func() lipgloss.Style {
		b := lipgloss.DoubleBorder()
		b.Right = "├"
		b.Left = "┤"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 2).BorderForeground(lipgloss.Color("39"))
	}()

	TitleStyleInactive = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		b.Left = "┤"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 2).BorderForeground(lipgloss.Color("255"))
	}()

	TitleStyleDisabled = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		b.Left = "┤"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 2).BorderForeground(lipgloss.Color("240")).Foreground(lipgloss.Color("240"))
	}()
)
