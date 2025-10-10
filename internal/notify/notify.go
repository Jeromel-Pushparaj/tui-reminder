package notify

import (
	"os/exec"
	"fmt"
)

func SendNotification(title, message string) error{
	cmd := exec.Command("notify-send", title, message)
	if err:=cmd.Run(); err != nil {
		return fmt.Errorf("failed to send notification: %v", err)
	}
	return nil

}
