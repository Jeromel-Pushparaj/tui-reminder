package reminder

import (
    "fmt"
    "time"
    "os/exec"
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
                sendNotification(r.Message)
            }
        case <-stopChan:
            fmt.Printf("[Reminder %d] Stopped.\n", r.ID)
            return
        }
    }
}

func sendNotification(msg string) {
    // Works on Linux systems
    cmd := exec.Command("notify-send", msg)
    _ = cmd.Run()
}
