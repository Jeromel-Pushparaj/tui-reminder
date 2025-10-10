package main

import (
	"github.com/Jeromel-Pushparaj/tui-reminder/internal/reminder"
)

func main() {
	reminders, _ := reminder.LoadReminders()
	stopChans := make(map[int]chan bool)

	// Start all active reminders
	for i := range reminders {
		if reminders[i].Active {
			stopChans[reminders[i].ID] = make(chan bool)
			go reminders[i].Start(stopChans[reminders[i].ID])
		}
	}
	select {}
}
