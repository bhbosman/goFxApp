package DateTimeService

import (
	"time"
)

type DateTimeService struct {
}

func (self *DateTimeService) Now() time.Time {
	return time.Now()
}

func NewDateTimeService() *DateTimeService {
	return &DateTimeService{}
}
