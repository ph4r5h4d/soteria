package models

type LogInterface interface {
	Build() (LogInterface, error)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Debug(msg string)
}
