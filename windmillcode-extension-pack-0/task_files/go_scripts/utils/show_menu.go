package utils

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)


type ShowMenuModel struct {
	Prompt string
	Choices  []string           // items on the to-do list
	cursor   int                // which to-do list item our cursor is pointing at
	Selected map[int]string  // which to-do items are selected
}



func ShowMenu( cliInfo ShowMenuModel, enableOtherOption  interface{} ) string  {
	cliInfo.Selected = make(map[int]string)
	p := tea.NewProgram(cliInfo)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	var value string
	for _, v := range cliInfo.Selected {
			value = v
			break
	}
	return value
}




func (m ShowMenuModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m ShowMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

			// Cool, what was the actual key pressed?
			switch msg.String() {

			// These keys should exit the program.
			case "ctrl+c", "q":
					return m, tea.Quit

			// The "up" and "k" keys move the cursor up
			case "up", "k":
					if m.cursor > 0 {
							m.cursor--
					}

			// The "down" and "j" keys move the cursor down
			case "down", "j":
					if m.cursor < len(m.Choices)-1 {
							m.cursor++
					}

			// The "enter" key and the spacebar (a literal space) toggle
			// the selected state for the item that the cursor is pointing at.
			case "enter", " ":
					_, ok := m.Selected[m.cursor]
					for key := range m.Selected {
						delete(m.Selected, key)
					}
					if ok {
							delete(m.Selected, m.cursor)
					} else {
							m.Selected[m.cursor] = m.Choices[m.cursor]
					}
					return m, tea.Quit
			}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m ShowMenuModel) View() string {
	// The header
	s := m.Prompt +"\n"

	// Iterate over our choices
	for i, choice := range m.Choices {

			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
					cursor = ">" // cursor!
			}

			// Is this choice selected?
			checked := " " // not selected
			if _, ok := m.Selected[i]; ok {
					checked = "x" // selected!
			}

			// Render the row
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	// s += "\nPress q to confirm selection.\n"

	// Send the UI for rendering
	return s
}
