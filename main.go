package main

import (
	"fmt"

	"github.com/9elements/u-root-menu/internal/terminal"
	"github.com/9elements/u-root-menu/pkg/config"
	tea "github.com/charmbracelet/bubbletea"
)

var Version = "0.0.0-dev"

func main() {
	fmt.Println("Version:", Version)
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	fmt.Println("Config loaded:", cfg)

	terminal := terminal.Setup()
	if _, err := tea.NewProgram(terminal).Run(); err != nil {
		fmt.Printf("Error running terminal: %v\nDropping into shell.", err)
		return
	}
}
