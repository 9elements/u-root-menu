package terminal

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Tab  key.Binding
	Quit key.Binding
}

var keys = func() keyMap {
	return keyMap{
		Tab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "Switch tabs"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
		),
	}
}()
