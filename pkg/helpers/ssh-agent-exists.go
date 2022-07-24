package helpers

import (
	"os/exec"
)

func CheckIfSshAgentExists() bool {
	_, err := exec.LookPath("ssh-add")
	_, err = exec.LookPath("ssh-agent")
	if err != nil {
		return false
	}
	return true
}
