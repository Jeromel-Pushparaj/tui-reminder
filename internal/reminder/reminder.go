/*
Package reminder
it's responsible for seting the reminder routine and organise the reminder from the json.
*/
package reminder

import (
	"fmt"
	"time"

	"github.com/Jeromel-Pushparaj/tui-reminder/internal/notify"
)

type Reminder struct {
	ID       int           `json:"id"`
	Message  string        `json:"message"`
	Interval time.Duration `json:"interval"`
	Active   bool          `json:"active"`
}

// Start runs a goroutine that sends a notification periodically
func (r *Reminder) Start(stopChan chan bool) {
	fmt.Printf("[Reminder %d] Started: %s every %v\n", r.ID, r.Message, r.Interval)
	ticker := time.NewTicker(r.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if r.Active {
				err := notify.SendNotification("Mom is notifying", r.Message)
				if err != nil {
					fmt.Printf("%s", err)
				}
			}
		case <-stopChan:
			fmt.Printf("[Reminder %d] Stopped.\n", r.ID)
			return
		}
	}
}
