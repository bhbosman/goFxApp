package multiLogger

import (
	"context"
	"github.com/bhbosman/gocommon/Services/IFxService"
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

type LoggerFileService struct {
	list     []io.Closer
	interval time.Duration
	state    IFxService.State
}

func (self *LoggerFileService) State() IFxService.State {
	return self.state
}

func (self *LoggerFileService) ServiceName() string {
	return "LoggerFileService"
}

func NewLoggerFileService(interval time.Duration) *LoggerFileService {
	return &LoggerFileService{
		list:     nil,
		interval: interval,
	}
}

func (self *LoggerFileService) OnStart(ctx context.Context) error {
	self.state = IFxService.Started
	return nil
}

func (self *LoggerFileService) OnStop(ctx context.Context) error {
	self.state = IFxService.Stopped
	return self.Close()
}

func (self *LoggerFileService) Close() error {
	var err error
	for _, closer := range self.list {
		err = multierr.Append(err, closer.Close())
	}
	return err
}
func (self *LoggerFileService) Build(
	cfg zapcore.EncoderConfig,
	resourceUrl string,
	enabler zapcore.LevelEnabler) (zapcore.Core, error) {
	f, err := NewLoggerFile(self.interval, resourceUrl)
	if err != nil {
		return nil, err
	}
	self.list = append(self.list, f)
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		f,
		enabler), nil
}
