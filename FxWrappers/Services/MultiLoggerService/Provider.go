package multiLogger

import (
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/url"
	"os"
	"time"
)

func InvokeMultiLogFileService() fx.Option {
	return fx.Options()
}
func ProvideMultiLogFileService() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Target: func(params struct {
					fx.In
				}) (*LoggerFileService, error) {
					instance := NewLoggerFileService(time.Hour)
					return instance, nil
				},
			}),
		fx.Provide(
			fx.Annotated{
				Group: "ZapCore.Core.Loggers",
				Target: func(params struct {
					fx.In
					MultiLogFileService *LoggerFileService
					Config              *zap.Config
					ApplicationName     string `name:"ApplicationName" optional:"true"`
				}) (zapcore.Core, zapcore.Core, error) {
					dir, err := os.UserHomeDir()
					if err != nil {
						return nil, nil, err
					}
					if params.ApplicationName == "" {
						params.ApplicationName = "UnknownApplication"
					}

					loggerCore, err := params.MultiLogFileService.Build(
						params.Config.EncoderConfig,
						fmt.Sprintf("file:///%v/%v/log?interval=hour&file-prefix=LOG", dir, params.ApplicationName),
						zap.LevelEnablerFunc(
							func(level zapcore.Level) bool {
								return level >= zapcore.InfoLevel
							}))
					if err != nil {
						return nil, nil, err
					}
					errorCore, err := params.MultiLogFileService.Build(
						params.Config.EncoderConfig,
						fmt.Sprintf("file:///%v/%v/log?interval=hour&file-prefix=ERR", dir, params.ApplicationName),
						zap.LevelEnablerFunc(
							func(level zapcore.Level) bool {
								return level >= zapcore.ErrorLevel
							}))
					if err != nil {
						return nil, nil, err
					}
					return loggerCore, errorCore, nil

				},
			}),
		fx.Provide(
			fx.Annotated{
				Group: "ZapCore.Core.Errors",
				Target: func(params struct {
					fx.In
					Config          *zap.Config
					ApplicationName string `name:"ApplicationName" optional:"true"`
				}) (zapcore.WriteSyncer, error) {
					dir, err := os.UserHomeDir()
					if err != nil {
						return nil, err
					}
					if params.ApplicationName == "" {
						params.ApplicationName = "UnknownApplication"
					}
					parseUrl, err := url.Parse(fmt.Sprintf("file:///%v/%v/error?interval=hour", dir, params.ApplicationName))
					if err != nil {
						return nil, err
					}
					errorFile := NewErrorFile(time.Hour, parseUrl, parseUrl.Path[1:])
					err = os.MkdirAll(errorFile.folder, os.ModeDir)

					return errorFile, nil
				},
			}),
		fx.Invoke(
			func(params struct {
				fx.In
				Lifecycle         fx.Lifecycle
				LoggerFileService *LoggerFileService
			}) error {
				params.Lifecycle.Append(fx.Hook{OnStart: params.LoggerFileService.OnStart, OnStop: params.LoggerFileService.OnStop})
				return nil
			}),
	)
}
