package main

import (
	"fmt"
	"time"
	"github.com/Jeromel-Pushparaj/tui-reminder/internal/reminder"
)
func main() {
    reminders, _:= reminder.LoadReminders()
    stopChans := make(map[int]chan bool)

    // Start all active reminders
    for i := range reminders {
        if reminders[i].Active {
            stopChans[reminders[i].ID] = make(chan bool)
            go reminders[i].Start(stopChans[reminders[i].ID])
        }
    }


	for{
        fmt.Println("\nCommands: add, list, stop, exit")
        var cmd string
        fmt.Scan(&cmd)

        switch cmd {
        case "add":
            var msg string
            var interval int
            fmt.Print("Message: ")
            fmt.Scanln(&msg)
            fmt.Print("Interval (min): ")
            fmt.Scanln(&interval)
			
    		newReminder := reminder.Reminder{ 
				ID:       len(reminders) + 1,
        		Message:  msg,
        		Interval: time.Duration(interval) * time.Minute,
        		Active:   true,
    		}
			reminders = append(reminders, newReminder) 
			reminder.SaveReminders(reminders)
            ch := make(chan bool)
            stopChans[newReminder.ID] = ch
            go newReminder.Start(ch)

        case "list":
            for _, r := range reminders {
                fmt.Printf("[%d] %s | every %v | Active: %v\n",
                    r.ID, r.Message, r.Interval, r.Active)
            }

        case "stop":
            var id int
            fmt.Print("Enter reminder ID to stop: ")
            fmt.Scan(&id)
            if ch, ok := stopChans[id]; ok {
                ch <- true
                delete(stopChans, id)
                for i := range reminders {
                    if reminders[i].ID == id {
                        reminders[i].Active = false
                    }
                }
                reminder.SaveReminders(reminders)
            }

        case "exit":
            fmt.Println("Exiting...")
            for _, ch := range stopChans {
                ch <- true
            }
            return
    }

	}
}
