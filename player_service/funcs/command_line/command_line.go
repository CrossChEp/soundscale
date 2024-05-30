package commandline

import (
	"bytes"
	"os/exec"
	"player_service/models"
)

func ShellOut(command string) models.CommandModel {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return models.CommandModel{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
		Error:  err,
	}
}
