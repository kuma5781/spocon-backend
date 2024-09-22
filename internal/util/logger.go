package util

var Logger *AppLogger

type AppLogger struct {
	logger GenericLogger
}

type GenericLogger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(err error, msg string)
}

func InitAppLogger(logger GenericLogger) {
	Logger = &AppLogger{logger: logger}
}

func (l *AppLogger) Debug(msg string) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Debug(msg)
}

func (l *AppLogger) Info(msg string) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Info(msg)
}

func (l *AppLogger) Warn(msg string) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Warn(msg)
}

func (l *AppLogger) Error(err error, msg string) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Error(err, msg)
}
