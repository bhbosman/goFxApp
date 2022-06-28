package goFxApp

import (
	"github.com/bhbosman/goConnectionManager"
	"github.com/bhbosman/goFxApp/Services/DateTimeService"
	"github.com/bhbosman/goFxApp/Services/MultiLoggerService"
	"github.com/bhbosman/goFxApp/internal"
	"github.com/bhbosman/goFxAppManager/FxServicesSlide"
	"github.com/bhbosman/goFxAppManager/Serivce"
	UiService2 "github.com/bhbosman/goUi/UiService"
	"github.com/bhbosman/goUi/UiSlides/GoFunctionCounterSlide"
	"github.com/bhbosman/goUi/UiSlides/cmSlide"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"go.uber.org/zap"

	"github.com/bhbosman/goUi/UiSlides/intoductionSlide"
	"github.com/bhbosman/gocommon/Services/Providers"
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

	provideTerminalApplicationOptions := fx.Options()
	invokeTerminalApplicationOptions := fx.Options()
	if !serviceApplication {
		provideTerminalApplicationOptions = fx.Options(
			fx.Populate(&terminalApplication, &logger),
			FxServicesSlide.ProvideServiceSlide(),
			intoductionSlide.ProvideCoverSlide(),
			cmSlide.ProvideConnectionManagerSlide(),
			UiService2.ProvideTerminalApplication(),
			GoFunctionCounterSlide.Provide(),
		)
		invokeTerminalApplicationOptions = fx.Options(
			cmSlide.InvokeConnectionManagerSlide(),
			UiService2.InvokeTerminalApplication(),
		)
	}

	fxOptions := fx2.NewFxApplicationOptions(
		time.Hour,
		time.Hour,
		fx.Options(
			GoFunctionCounter.Provide(),
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
			goConnectionManager.ProvideConnectionManager(),

			DateTimeService.ProvideDateTimeService(),
			multiLogger.ProvideMultiLogFileService(),
			Serivce.ProvideFxManager(),
			fx.Options(option...),
			multiLogger.InvokeMultiLogFileService(),
			internal.InvokeApps(),
			goConnectionManager.InvokeConnectionManager(),
			invokeTerminalApplicationOptions,
			Serivce.InvokeFxManager(),
			internal.InvokeApplicationContext(),
		),
	)
	fxApp := fx.New(fxOptions)

	return NewTerminalAppUsingFxApp(
		terminalApplication,
		shutDowner,
		fxApp,
		logger,
	)
}
