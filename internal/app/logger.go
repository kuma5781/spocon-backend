package app

import (
	"fmt"
	"io"
	"log/slog"
	"runtime/debug"
)

type LogConfig struct {
	Name   string
	Level  string
	Out    io.Writer
	OutErr io.Writer
}

func NewLogConfig(name string, out io.Writer, outErr io.Writer) *LogConfig {
	return &LogConfig{
		Name:   name,
		Level:  "debug", // TODO: 引数でConfigを受け取って、環境ごとに設定したい
		Out:    out,
		OutErr: outErr,
	}
}

func InitLogger(out io.Writer, outErr io.Writer) *SLogLogger {
	logConfig := NewLogConfig("app_log", out, outErr)
	return newSlogger(logConfig)
}

type SLogLogger struct {
	outLogger    *slog.Logger
	outErrLogger *slog.Logger
}

func (s *SLogLogger) Debug(msg string) {
	s.outLogger.Debug(msg)
}

func (s *SLogLogger) Info(msg string) {
	s.outLogger.Info(msg)
}

func (s *SLogLogger) Warn(msg string) {
	s.outLogger.Warn(msg)
}

func (s *SLogLogger) Error(err error, msg string) {
	combinedMsg := fmt.Sprintf("%s: %v", msg, err)
	s.outLogger.Error(combinedMsg, "stacktrace", debug.Stack())
}

func newSlogger(conf *LogConfig) *SLogLogger {
	var level slog.Level
	switch conf.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	sLogger := &SLogLogger{}

	opts := &slog.HandlerOptions{
		Level: level,
	}
	logger := slog.New(slog.NewJSONHandler(conf.Out, opts))
	sLogger.outLogger = logger

	errLogger := slog.New(slog.NewJSONHandler(conf.OutErr, opts))
	sLogger.outErrLogger = errLogger

	return sLogger
}
