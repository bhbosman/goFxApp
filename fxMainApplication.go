package goFxApp

import (
	"github.com/bhbosman/goFxApp/Services/DateTimeService"
	"github.com/bhbosman/goFxApp/Services/MultiLoggerService"
	"github.com/bhbosman/goFxApp/internal"
	"github.com/bhbosman/goFxAppManager/FxServicesSlide"
	"github.com/bhbosman/goFxAppManager/Serivce"
	UiService2 "github.com/bhbosman/goUi/UiService"
	"github.com/bhbosman/gocommon/Services/Providers"
	"github.com/bhbosman/gocommon/Services/implementations/ConnectionManagerService"
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

	provideTerminalApplicationOptions := fx.Options()
	invokeTerminalApplicationOptions := fx.Options()
	if !serviceApplication {
		provideTerminalApplicationOptions = fx.Options(
			fx.Populate(&terminalApplication),
			FxServicesSlide.Dddddd(),
			UiService2.ProvideTerminalApplication())
		invokeTerminalApplicationOptions = fx.Options(
			UiService2.InvokeTerminalApplication())
	}

	fxOptions := fx2.NewFxApplicationOptions(
		time.Hour,
		time.Hour,
		fx.Options(
			fx.Populate(&shutDowner),
			provideTerminalApplicationOptions,
			fx.Provide(
				fx.Annotated{
					Target: func() string {
						return applicationName
					},
					Name: "ApplicationName",
				}),
			internal.ProvideApplicationContext(),
			Providers.ProvideUniqueSessionNumber(),
			Providers.ProvideNewUniqueReferenceService(),
			PubSub.ProvidePubSub("Application"),
			Zap2.ProvideZapCoreEncoderConfigForDev(),
			Zap2.ProvideZapCoreEncoderConfigForProd(),
			Zap2.ProvideZapConfigForDev(nil, nil),
			Zap2.ProvideZapConfigForProd(nil, nil),
			Zap2.ProvideZapLogger(),
			ConnectionManagerService.ProvideConnectionManager(),
			internal.ProvideFxWithLogger(),
			DateTimeService.ProvideDateTimeService(),
			multiLogger.ProvideMultiLogFileService(),
			Serivce.ProvideFxManager(),
			fx.Options(option...),
			internal.InvokeApplicationContext(),
			multiLogger.InvokeMultiLogFileService(),
			internal.InvokeApps(),
			invokeTerminalApplicationOptions,
			ConnectionManagerService.InvokeConnectionManager(),
			Serivce.InvokeFxManager(),
		),
	)
	fxApp := fx.New(fxOptions)

	return NewTerminalAppUsingFxApp(
		terminalApplication,
		shutDowner,
		fxApp,
	)
}
