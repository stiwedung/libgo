package log

import (
	"bufio"
	"io"
)

const (
	timeFormat       = "2006-01-02 15:04:05"
	maxOuputBuffSize = 8192
)

type IPrinter interface {
	Close() error
	write(*log)
	flush()
}

func newPrinter(output io.WriteCloser, isTerminal bool) IPrinter {
	if isTerminal {
		return &terminalPrinter{
			Closer: output,
			output: bufio.NewWriter(output),
		}
	}
	return &filePrinter{
		Closer: output,
		output: bufio.NewWriter(output),
	}
}

type terminalPrinter struct {
	io.Closer
	output *bufio.Writer
}

func (p *terminalPrinter) write(log *log) {
	p.output.WriteString(log.lvl.colorfulText())
	p.output.WriteByte(' ')
	p.output.WriteString(log.time.Format(timeFormat))
	p.output.WriteByte(' ')
	if log.caller != "" {
		p.output.WriteString(log.caller)
		p.output.WriteByte(' ')
	}
	p.output.WriteString(log.msg)
	if log.newline {
		p.output.WriteByte('\n')
	}
	p.output.Flush()
}

func (p *terminalPrinter) flush() {}

type filePrinter struct {
	io.Closer
	output *bufio.Writer
}

func (p *filePrinter) write(log *log) {
	p.output.WriteString(log.lvl.rawText())
	p.output.WriteByte(' ')
	p.output.WriteString(log.time.Format(timeFormat))
	p.output.WriteByte(' ')
	if log.caller != "" {
		p.output.WriteString(log.caller)
		p.output.WriteByte(' ')
	}
	p.output.WriteString(log.msg)
	if log.newline {
		p.output.WriteByte('\n')
	}
	if p.output.Size() >= maxOuputBuffSize {
		p.output.Flush()
	}
}

func (p *filePrinter) flush() {
	p.output.Flush()
}

type multiPrinter struct {
	writers []IPrinter
}

func (p *multiPrinter) Close() error {
	for _, w := range p.writers {
		w.Close()
	}
	return nil
}

func (p *multiPrinter) write(log *log) {
	for _, w := range p.writers {
		w.write(log)
	}
}

func (p *multiPrinter) flush() {
	for _, w := range p.writers {
		w.flush()
	}
}
