package boot

import (
	"fmt"
	"log"
	"strings"

	"github.com/alimsk/list"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/u-root/u-root/pkg/boot"
	"github.com/u-root/u-root/pkg/boot/localboot"
	"github.com/u-root/u-root/pkg/mount"
	"github.com/u-root/u-root/pkg/mount/block"
	"github.com/u-root/u-root/pkg/ulog"
)

type Model struct {
	Viewport *viewport.Model
	Help     help.Model

	// Menu menu.BIOSMenu
	list   list.Model
	images []boot.OSImage

	keyMap keyMap
}

func Setup() *Model {
	blockDevs, err := block.GetBlockDevices()
	if err != nil {
		log.Fatal("No available block devices to boot from")
	}

	fmt.Printf("Found %d block devices\n", len(blockDevs))

	for _, dev := range blockDevs {
		fmt.Printf("Device: %s\n", dev)
		size, err := dev.Size()
		if err != nil {
			fmt.Printf("Error getting size: %v\n", err)
		}
		fmt.Printf("Size: %d\n", size)
	}

	// Try to only boot from "good" block devices.
	blockDevs = blockDevs.FilterZeroSize()

	fmt.Printf("Found %d block devices (after filtering)\n", len(blockDevs))

	l := ulog.Null

	mountPool := &mount.Pool{}
	images, err := localboot.Localboot(l, blockDevs, mountPool)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d images\n", len(images))
	for _, img := range images {
		fmt.Printf("Image: %+v\n", img.Label())
	}

	simpleList := list.NewSimpleAdapter(list.SimpleItemList{})
	for _, img := range images {
		simpleList.Append(list.SimpleItem{
			Title:    img.Label(),
			Desc:     img.Label(),
			Disabled: false,

			Options:        []string{},
			SelectedOption: "",
		})
	}

	return &Model{
		list:   list.New(simpleList),
		images: images,

		keyMap: keys,
	}
}

func (m *Model) Init() tea.Cmd {
	m.list.Focus()
	return nil
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
		case key.Matches(msg, m.keyMap.Enter):
			fmt.Printf("Selected: %s\n", m.images[m.list.ItemFocus()].Label())
			adapter := m.list.Adapter.(*list.SimpleAdapter)
			item := adapter.FilteredItemAt(m.list.ItemFocus())
			if item.Disabled {
				break
			}
			if item.Title == m.images[m.list.ItemFocus()].Label() {
				fmt.Printf("Loading Image: %s\n", m.images[m.list.ItemFocus()].Label())
				err := m.images[m.list.ItemFocus()].Load()
				if err != nil {
					fmt.Printf("Error loading image: %v", err)
					break
				}

				fmt.Printf("Executing Image: %s\n", m.images[m.list.ItemFocus()].Label())
				err = boot.Execute()
				if err != nil {
					fmt.Printf("Error executing image: %v", err)
					break
				}
			} else {
				fmt.Printf("Selected item: %s mismatch %s\n", item.Title, m.images[m.list.ItemFocus()].Label())
			}
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	menuDoc := strings.Builder{}

	ws := lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("39")).
		PaddingLeft(20).
		PaddingRight(20).
		PaddingTop(2).
		Align(lipgloss.Left).
		Border(lipgloss.RoundedBorder()).
		Width(m.Viewport.Width - 7)

	// Render the menu
	menuItems := m.list.View()
	menuDoc.WriteString(menuItems)

	docHeight := strings.Count(menuDoc.String(), "\n")
	requiredNewlinesForPadding := m.Viewport.Height - docHeight - 13

	menuDoc.WriteString(strings.Repeat("\n", max(0, requiredNewlinesForPadding)))

	return ws.Render(menuDoc.String())
}
