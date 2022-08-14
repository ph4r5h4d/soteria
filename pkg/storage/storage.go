package storage

import (
	"errors"
	"github.com/ph4r5h4d/soteria/models"
	"github.com/ph4r5h4d/soteria/pkg/storage/git"
)

var storage = map[string]interface{}{
	"git": &git.Git{},
}

func Build(name string, config models.Config, logger models.LogInterface) (models.StorageInterface, error) {
	storageInstance, ok := storage[name]
	if !ok {
		return nil, errors.New("unsupported storage")
	}

	storage, err := storageInstance.(models.StorageInterface).Build(config, logger)
	if err != nil {
		return nil, err
	}
	return storage, nil
}
