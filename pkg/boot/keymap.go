package boot

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Enter key.Binding
	Quit  key.Binding
}

var keys = func() keyMap {
	return keyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
		),
		Enter: key.NewBinding(
			key.WithKeys("b"),
		),
	}
}()
