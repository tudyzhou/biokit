package utils

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// readline, iterable
func Readline(reader *bufio.Reader) (string, error) {
	line, isPrefix, err := reader.ReadLine()
	if !isPrefix {
		return string(line), err
	}
	buf := append([]byte(nil), line...)
	for isPrefix && err == nil {
		line, isPrefix, err = reader.ReadLine()
		buf = append(buf, line...)
	}
	return string(buf), err
}

// read whole file to lines slice
func ReadLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)

	if file, err = os.Open(path); err != nil {
		return
	}

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))

	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}
