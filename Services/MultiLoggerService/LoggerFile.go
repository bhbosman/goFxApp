package multiLogger

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"io"
	"net/url"
	"os"
	"time"
)

type LoggerFile struct {
	BaseFile
	currentFile io.WriteCloser
	errorState  error
	closed      bool
}

func NewLoggerFile(interval time.Duration, urlResource string) (
	interface {
		io.Closer
		zapcore.WriteSyncer
	}, error) {
	urlPath, err := url.Parse(urlResource)
	if err != nil {
		return nil, err
	}
	if urlPath.Scheme != "file" {
		return nil, fmt.Errorf("could nt support %v. Full path: %v", urlPath.Scheme, urlPath.String())
	}

	result := &LoggerFile{
		BaseFile: NewBaseFile(interval, urlPath, urlPath.Path[1:]),
	}
	err = os.MkdirAll(result.folder, os.ModeDir)

	if err != nil {
		return nil, err
	}
	err = result.calculateBucket()
	if err != nil {
		return nil, err
	}
	err = result.createFile()

	return result, nil
}

func (self *LoggerFile) Write(p []byte) (int, error) {
	if self.next.Sub(time.Now()) < 0 {
		err := self.calculateBucket()
		if err != nil {
			return 0, err
		}
		err = self.createFile()
		if err != nil {
			return 0, err
		}
	}

	if self.closed {
		return 0, io.EOF
	}
	if self.errorState != nil {
		return 0, self.errorState
	}
	if self.currentFile == nil {
		return 0, fmt.Errorf("currentFile equal to nil")
	}

	return self.currentFile.Write(p)
}

func (self *LoggerFile) Sync() error {
	return nil
}

func (self *LoggerFile) Close() error {
	self.closed = true
	temp := self.currentFile
	self.currentFile = nil
	if temp != nil {
		err := temp.Close()
		if err != nil {
			self.errorState = err
		} else {
			self.errorState = io.EOF
		}
	}
	return self.errorState
}

func (self *LoggerFile) createFile() error {
	query, err := url.ParseQuery(self.url.RawQuery)
	if err != nil {
		return err
	}

	prefix := query.Get("file-prefix")
	if prefix == "" {
		prefix = "logger"
	}
	fileName := self.BuildFileName(prefix)
	if self.closed {
		return io.EOF
	}

	if self.currentFile != nil {
		err := self.currentFile.Close()
		if err != nil {
			return err
		}
		self.currentFile = nil
	}
	self.currentFile, err = self.openFile(fileName)
	if err != nil {
		return err
	}
	return nil
}
