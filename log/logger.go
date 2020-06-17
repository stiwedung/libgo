package log

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/onsi/ginkgo/reporters/stenographer/support/go-isatty"
)

type log struct {
	newline bool
	lvl     Level
	time    time.Time
	msg     string
	caller  string
}

func NewLogger(file *os.File, opts ...Option) *Logger {
	logger := &Logger{}
	logger.pool.New = func() interface{} {
		return &log{}
	}
	logger.inputQ = make(chan *log, 8192)
	logger.closed = make(chan struct{})
	isTerminal := isatty.IsTerminal(file.Fd())
	logger.printer = newPrinter(file, isTerminal)
	logger.newline = true
	logger.skip = 2
	for _, opt := range opts {
		opt(logger)
	}
	logger.wg.Add(1)
	go logger.run()
	return logger
}

type Logger struct {
	pool    sync.Pool
	printer IPrinter
	newline bool
	caller  bool
	level   Level
	inputQ  chan *log
	closed  chan struct{}
	wg      sync.WaitGroup
	skip    int
}

func (logger *Logger) requireLog() *log {
	return logger.pool.Get().(*log)
}

func (logger *Logger) releaseLog(l *log) {
	logger.pool.Put(l)
}

func (logger *Logger) run() {
	ticker := time.NewTicker(time.Second)
	defer logger.close()
	for {
		select {
		case log := <-logger.inputQ:
			logger.printer.write(log)
			if log.lvl == FatalLevel {
				logger.printer.flush()
				os.Exit(1)
			}
		case <-ticker.C:
			logger.printer.flush()
		case <-logger.closed:
			return
		}
	}
}

func (logger *Logger) log(lvl Level, msg string) {
	log := logger.requireLog()
	log.lvl = lvl
	log.newline = logger.newline
	log.time = time.Now()
	log.msg = msg
	if logger.caller {
		_, file, line, _ := runtime.Caller(logger.skip)
		log.caller = fmt.Sprintf("%s:%d", file, line)
	} else {
		log.caller = ""
	}
	logger.inputQ <- log
}

func (logger *Logger) SetLevel(name string) {
	level := ParseLevel(name)
	logger.level = level
}

func (logger *Logger) AddPrinter(file *os.File) {
	isTerminal := isatty.IsTerminal(file.Fd())
	printer := newPrinter(file, isTerminal)
	if p, ok := logger.printer.(*multiPrinter); ok {
		p.writers = append(p.writers, printer)
	} else {
		mp := &multiPrinter{}
		mp.writers = append(mp.writers, logger.printer)
		mp.writers = append(mp.writers, printer)
		logger.printer = mp
	}
}

func (logger *Logger) Close() {
	logger.level = DisableLevel
	close(logger.closed)
	logger.wg.Wait()
}

func (logger *Logger) close() {
outter:
	for {
		select {
		case log := <-logger.inputQ:
			logger.printer.write(log)
		default:
			break outter
		}
	}
	logger.printer.flush()
	logger.printer.Close()
	logger.wg.Done()
}

func (logger *Logger) Debug(args ...interface{}) {
	if logger.level > DebugLevel {
		return
	}
	logger.log(DebugLevel, fmt.Sprint(args...))
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	if logger.level > DebugLevel {
		return
	}
	msg := fmt.Sprintf(format, args...)
	logger.log(DebugLevel, msg)
}

func (logger *Logger) Info(args ...interface{}) {
	if logger.level > InfoLevel {
		return
	}
	logger.log(InfoLevel, fmt.Sprint(args...))
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	if logger.level > InfoLevel {
		return
	}
	msg := fmt.Sprintf(format, args...)
	logger.log(InfoLevel, msg)
}

func (logger *Logger) Warn(args ...interface{}) {
	if logger.level > WarnLevel {
		return
	}
	logger.log(WarnLevel, fmt.Sprint(args...))
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	if logger.level > WarnLevel {
		return
	}
	msg := fmt.Sprintf(format, args...)
	logger.log(WarnLevel, msg)
}

func (logger *Logger) Error(args ...interface{}) {
	if logger.level > ErrorLevel {
		return
	}
	logger.log(ErrorLevel, fmt.Sprint(args...))
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	if logger.level > ErrorLevel {
		return
	}
	msg := fmt.Sprintf(format, args...)
	logger.log(ErrorLevel, msg)
}

func (logger *Logger) Fatal(args ...interface{}) {
	if logger.level > FatalLevel {
		return
	}
	logger.log(FatalLevel, fmt.Sprint(args...))
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	if logger.level > FatalLevel {
		return
	}
	msg := fmt.Sprintf(format, args...)
	logger.log(FatalLevel, msg)
}
