package reminder

import (
	"fmt"
    "encoding/json"
	"path/filepath"
    "os"
)

var filePath = filepath.Join("internal", "data", "reminders.json")
//SaveReminders writes all reminders to JSON file
func SaveReminders(reminders []Reminder) error {
    data, err := json.MarshalIndent(reminders, "", "  ")
    if err != nil {
        return err
    }
	fmt.Println("coming inside....")
    return os.WriteFile(filePath, data, 0644)
}

// LoadReminders reads all reminders from JSON file
func LoadReminders() ([]Reminder, error) {
    file, err := os.ReadFile(filePath)
    if err != nil {
        if os.IsNotExist(err) {
            return []Reminder{}, nil // empty if file not exists
        }
        return nil, err
    }

    var reminders []Reminder
    err = json.Unmarshal(file, &reminders)
    if err != nil {
        return nil, err
    }
    return reminders, nil
}
