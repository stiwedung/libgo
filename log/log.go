package log

import (
	"os"
)

var std = NewLogger(os.Stdout, WriteCallerOption(true), LevelOption(DebugLevel), SkipOption(3))

func GetLogger() *Logger {
	return std
}

func SetLevel(name string) {
	std.SetLevel(name)
}

func Close() {
	std.Close()
}

func Debug(msg string) {
	std.Debug(msg)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func Info(msg string) {
	std.Info(msg)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Warn(msg string) {
	std.Warn(msg)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Error(msg string) {
	std.Error(msg)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Fatal(msg string) {
	std.Fatal(msg)
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}
