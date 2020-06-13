package log

type Option func(logger *Logger)

func WriteCallerOption(val bool) Option {
	return func(logger *Logger) {
		logger.caller = val
	}
}

func NewLineOption(val bool) Option {
	return func(logger *Logger) {
		logger.newline = val
	}
}

func LevelOption(val Level) Option {
	return func(logger *Logger) {
		logger.level = val
	}
}

func SkipOption(val int) Option {
	return func(logger *Logger) {
		logger.skip = val
	}
}
