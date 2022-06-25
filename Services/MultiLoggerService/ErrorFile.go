package multiLogger

import (
	"net/url"
	"os"
	"time"
)

type ErrorFile struct {
	BaseFile
}

func NewErrorFile(interval time.Duration, url *url.URL, folder string) *ErrorFile {
	return &ErrorFile{
		BaseFile: NewBaseFile(interval, url, folder),
	}
}

func (self *ErrorFile) Write(p []byte) (n int, err error) {
	err = self.calculateBucket()
	if err != nil {
		return 0, err
	}
	fileName := self.BuildFileName("error")
	openFile, err := self.openFile(fileName)
	if err != nil {
		return 0, err
	}
	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {

		}
	}(openFile)
	return openFile.Write(p)
}

func (self *ErrorFile) Sync() error {
	return nil
}
