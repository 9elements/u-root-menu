package main

import (
	"fmt"

	"github.com/9elements/u-root-menu/internal/terminal"
	"github.com/9elements/u-root-menu/pkg/config"
	"github.com/9elements/u-root-menu/pkg/menu"
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

	// XXX: This list needs to be parsed from Firmware.
	// We might need a more complex structure to represent the menu.
	terminal := terminal.Setup(Version, menu.BIOSMenu{
		Name:     "BIOS",
		HelpText: "BIOS Configuration",
		Options:  []menu.BIOSOption{},
		SubMenus: []menu.BIOSMenu{},
	})
	if _, err := tea.NewProgram(terminal).Run(); err != nil {
		fmt.Printf("Error running terminal: %v\nDropping into shell.", err)
		return
	}
}
