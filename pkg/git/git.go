package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/ph4r5h4d/soteria/models"
	"github.com/ph4r5h4d/soteria/pkg/helpers"
	"os"
)

func Clone(logger models.LogInterface, repository string) error {
	home, _ := os.UserHomeDir()
	gitDir := home + "/.soteria/storage"
	if !helpers.CheckIfSshAgentExists() {
		logger.Error("SSH Agent is not installed")
	}

	auth, err := setupAuth()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = git.PlainClone(gitDir, false, &git.CloneOptions{
		URL:      repository,
		Progress: os.Stdout,
		Auth:     auth,
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func setupAuth() (*ssh.PublicKeysCallback, error) {
	sshAuth, err := ssh.NewSSHAgentAuth("git")
	if err != nil {
		return nil, err
	}
	return sshAuth, nil
}
