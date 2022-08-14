package git

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/ph4r5h4d/soteria/models"
	"github.com/ph4r5h4d/soteria/pkg/helpers"
	"os"
)

type Git struct {
	logger models.LogInterface
	config models.Config
}

func (g Git) Build(config models.Config, logger models.LogInterface) (models.StorageInterface, error) {
	g.config = config
	g.logger = logger
	return g, nil
}

func (g Git) Init() error {
	home, _ := os.UserHomeDir()
	gitDir := home + "/.soteria/storage"
	if !helpers.CheckIfSshAgentExists() {
		return errors.New("SSH Agent is not installed")
	}

	auth, err := setupAuth()
	if err != nil {
		return err
	}

	_, err = git.PlainClone(gitDir, false, &git.CloneOptions{
		URL:      g.config.Git.Repository,
		Progress: os.Stdout,
		Auth:     auth,
	})

	if err != nil {
		return err
	}
	return nil
}

func (g Git) Sync() error {
	return nil
}

func setupAuth() (*ssh.PublicKeysCallback, error) {
	sshAuth, err := ssh.NewSSHAgentAuth("git")
	if err != nil {
		return nil, err
	}
	return sshAuth, nil
}
