package logger

// Logger defines a minimal logger.
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

type LeveledLogger interface {
	Logger

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}
