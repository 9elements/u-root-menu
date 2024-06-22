package dashboard

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Quit key.Binding
}

var keys = func() keyMap {
	return keyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
		),
	}
}()
