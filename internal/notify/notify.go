package notify

import (
	"fmt"
	"os/exec"
)

func SendNotification(title, message string) error {
	cmd := exec.Command("notify-send", title, message)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to send notification: %v", err)
	}
	// soundPath, _ := filepath.Abs("/home/jeromel/Public/linuxplayground/tui-reminder/internal/notify/notify_sound.wav")
	// cmdAplay := exec.Command("aplay", soundPath) // Replace with your sound file path
	// err := cmdAplay.Run()
	// if err != nil {
	// 	return fmt.Errorf("error playing sound with aplay: %v", err)
	// }

	return nil

}
