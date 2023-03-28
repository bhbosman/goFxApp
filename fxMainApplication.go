package goFxApp

import (
	"fmt"
	"github.com/bhbosman/goConnectionManager"
	"github.com/bhbosman/goFxApp/Services/DateTimeService"
	"github.com/bhbosman/goFxApp/Services/MultiLoggerService"
	"github.com/bhbosman/goFxApp/Services/fileDumpService"
	"github.com/bhbosman/goFxApp/internal"
	"github.com/bhbosman/goFxAppManager/FxServicesSlide"
	"github.com/bhbosman/goFxAppManager/service"
	"github.com/bhbosman/goUi/UiService"
	"github.com/bhbosman/goUi/UiSlides/GoFunctionCounterSlide"
	"github.com/bhbosman/goUi/UiSlides/cmSlide"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/services/Providers"
	"github.com/cskr/pubsub"
	"go.uber.org/zap"
	"net/url"
	"os"

	"github.com/bhbosman/goUi/UiSlides/intoductionSlide"
	fx2 "github.com/bhbosman/gocommon/fx"

	"github.com/bhbosman/gocommon/fx/PubSub"
	Zap2 "github.com/bhbosman/gocommon/fx/Zap"
	"github.com/rivo/tview"
	"go.uber.org/fx"
	"time"
)

func NewFxMainApplicationServices(
	applicationName string,
	serviceApplication bool,
	option ...fx.Option,
) *TerminalAppUsingFxApp {
	var terminalApplication *tview.Application
	var shutDowner fx.Shutdowner
	var logger *zap.Logger

	ss := struct {
		fx.In
		PubSub *pubsub.PubSub `name:"Application"`
	}{}

	provideTerminalApplicationOptions := fx.Options()
	invokeTerminalApplicationOptions := fx.Options()
	if !serviceApplication {
		provideTerminalApplicationOptions = fx.Options(
			fx.Populate(&terminalApplication, &logger),
			FxServicesSlide.ProvideServiceSlide(), //--
			intoductionSlide.ProvideCoverSlide(),
			cmSlide.ProvideConnectionManagerSlide(),
			UiService.ProvideTerminalApplication(),
			GoFunctionCounterSlide.Provide(),
		)
		invokeTerminalApplicationOptions = fx.Options(
			cmSlide.InvokeConnectionManagerSlide(),
			UiService.InvokeTerminalApplication(),
		)
	}

	fxOptions := fx2.NewFxApplicationOptions(
		time.Hour,
		time.Hour,
		fx.Options(
			GoFunctionCounter.Provide(),
			fx.Populate(&shutDowner, &ss),
			provideTerminalApplicationOptions,
			fx.Provide(
				fx.Annotated{
					Name: "ApplicationName",
					Target: func(
						params struct {
							fx.In
						},
					) string {
						return applicationName
					},
				},
			),
			fx.Provide(
				fx.Annotated{
					Name: "ConfigurationFolder",
					Target: func(
						params struct {
							fx.In
							ApplicationName string `name:"ApplicationName"`
						},
					) (string, error) {
						dir, err := os.UserHomeDir()
						if err != nil {
							return "", err
						}
						folder := fmt.Sprintf("file:///%v/ApplicationLogFiles/%v/Configuration", dir, params.ApplicationName)
						urlPath, err := url.Parse(folder)
						if err != nil {
							return "", err
						}
						folder = urlPath.Path[1:]
						err = os.MkdirAll(folder, os.ModePerm)
						if err != nil {
							return "", err
						}
						return folder, nil
					},
				},
			),
			internal.ProvideApplicationContext(),
			Providers.ProvideUniqueSessionNumber(),
			Providers.ProvideNewUniqueReferenceService(),
			PubSub.ProvidePubSub("Application"),
			Zap2.ProvideZapCoreEncoderConfigForDev(),
			Zap2.ProvideZapCoreEncoderConfigForProd(),
			Zap2.ProvideZapConfigForDev(nil, nil),
			Zap2.ProvideZapConfigForProd(nil, nil),
			Zap2.ProvideZapLogger(),
			goConnectionManager.ProvideCreateConnectionManager(),
			DateTimeService.ProvideDateTimeService(),
			multiLogger.ProvideMultiLogFileService(),
			fileDumpService.Provide(),
			service.ProvideFxManager(), //---
			multiLogger.InvokeMultiLogFileService(),
			internal.InvokeApps(),
			goConnectionManager.InvokeConnectionManager(),
			invokeTerminalApplicationOptions,
			fx.Options(option...),
			internal.InvokeApplicationContext(),
			service.InvokeFxManagerStartAll(),
		),
	)
	fxApp := fx.New(fxOptions)

	return NewTerminalAppUsingFxApp(
		terminalApplication,
		shutDowner,
		fxApp,
		logger,
		ss.PubSub,
	)
}
