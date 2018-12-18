package logur

import (
	"bufio"
	"fmt"
	"io"
	"runtime"
)

// NewWriter creates a new writer from a Logger with a default Info level.
func NewWriter(logger Logger) *io.PipeWriter {
	return NewLevelWriter(logger, Info)
}

// NewLevelWriter creates a new writer from a Logger for a specific level of log events.
func NewLevelWriter(logger Logger, level Level) *io.PipeWriter {
	reader, writer := io.Pipe()

	go writerScanner(logger, level, reader)

	runtime.SetFinalizer(writer, writerFinalizer)

	return writer
}

func writerScanner(logger Logger, level Level, reader io.ReadCloser) {
	scanner := bufio.NewScanner(reader)

	logFunc := LevelFunc(logger, level)

	for scanner.Scan() {
		logFunc(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Error(fmt.Sprintf("error while reading from log pipe: %s", err))
	}

	_ = reader.Close()
}

func writerFinalizer(writer *io.PipeWriter) {
	_ = writer.Close()
}
