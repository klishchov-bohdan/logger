package logger

type Logger interface {
	Info() error
	Warning() error
	Error() error
	Debug() error
}
