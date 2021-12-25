package logger

type Logger interface {
	Info(str string)
	Warning(str string)
	Error(str string)
	Debug(str string)
}
