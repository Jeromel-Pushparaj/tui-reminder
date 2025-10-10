package ui

import (
    "fmt"
    "time"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/Jeromel-Pushparaj/tui-reminder/internal/reminder"
)

// Model holds the app state
type Model struct {
    reminders []string
    // cursor    int
    // choice    string
    quitting  bool
}

// Init runs when the program starts
func (m Model) Init() tea.Cmd {
    return nil
}

// Update handles user input and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    reminders, _:= reminder.LoadReminders()
    stopChans := make(map[int]chan bool)
    switch msg := msg.(type) {

    case tea.KeyMsg:
        switch msg.String() {

        case "ctrl+c", "q":
            m.quitting = true
            return m, tea.Quit

        case "a":
            var msg string
            var interval int
            fmt.Print("Message: ")
            fmt.Scanln(&msg)
            fmt.Print("Interval (sec): ")
            fmt.Scanln(&interval)
			
    		newReminder := reminder.Reminder{ 
				ID:       len(reminders) + 1,
        		Message:  msg,
        		Interval: time.Duration(interval) * time.Second,
        		Active:   true,
    		}
			reminders = append(reminders, newReminder) 
            m.reminders = append(m.reminders, msg)
			reminder.SaveReminders(reminders)
            ch := make(chan bool)
            stopChans[newReminder.ID] = ch
            go newReminder.Start(ch)

        case "d":
            if len(m.reminders) > 0 {
                m.reminders = m.reminders[:len(m.reminders)-1]
            }
        }
    }
    return m, nil
}

// View renders the UI
func (m Model) View() string {
    s := "Reminder TUI\n\n"
    s += "Reminders:\n"

    for _, r := range m.reminders {
        s += fmt.Sprintf("• %s\n", r)
    }

    s += "\n[a] Add [d] Delete [q] Quit\n"
    return s
}

// New creates the initial model
func New() Model {
    return Model{}
}

