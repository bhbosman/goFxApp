package multiLogger

import (
	"fmt"
	"net/url"
	"os"
	"time"
)

type BaseFile struct {
	interval time.Duration
	url      *url.URL
	folder   string
	current  time.Time
	next     time.Time
}

func (self *BaseFile) BuildFileName(prefix string) string {
	return fmt.Sprintf("%v\\%v_%v_%v.log",
		self.folder,
		prefix,
		self.current.Format("20060102_150405"),
		self.next.Format("20060102_150405"))
}

func (self *BaseFile) calculateBucket() error {
	query, err := url.ParseQuery(self.url.RawQuery)
	if err != nil {
		return err
	}
	var interval time.Duration
	switch query.Get("interval") {
	case "min":
		interval = time.Minute
	case "day":
		interval = time.Hour * 24
	case "hour":
		interval = time.Hour
	case "30m":
		interval = time.Minute * 30
	default:
		interval = self.interval
	}

	inSeconds := int64(interval) / int64(time.Second)
	d := time.Now().Unix() / inSeconds
	self.current = time.Unix(inSeconds*d, 0)
	self.next = self.current.Add(interval)
	return err
}

func (self *BaseFile) openFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
}

func NewBaseFile(interval time.Duration, url *url.URL, folder string) BaseFile {
	return BaseFile{
		interval: interval,
		url:      url,
		folder:   folder,
	}
}
