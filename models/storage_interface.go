package models

type StorageInterface interface {
	Build(config Config, logInterface LogInterface) (StorageInterface, error)
	Init() error
	Sync() error
}
