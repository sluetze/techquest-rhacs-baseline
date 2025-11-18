package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

var flagUUID string

func init() {
	parts := []string{
		"7b1f2b3e", "-4c5d-", "6e7f-", "8091-", "a2b3c4d5e6f7",
	}
	flagUUID = strings.Join(parts, "")
}

func main() {
	// Sleep for 70 minutes before starting
	time.Sleep(70 * time.Minute)

	for {
		cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo external-tool-invoked --flag=%s >/dev/null", flagUUID))
		if out, err := cmd.CombinedOutput(); err != nil {
			log.Printf("external command error: %v", err)
		} else if len(out) > 0 {
			_ = out
		}
		time.Sleep(5 * time.Minute)
	}
}
