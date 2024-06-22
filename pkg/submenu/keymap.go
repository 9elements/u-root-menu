package submenu

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Enter key.Binding
	Quit  key.Binding
}

var keys = func() keyMap {
	return keyMap{
		Enter: key.NewBinding(
			key.WithKeys("enter"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q"),
		),
	}
}()
