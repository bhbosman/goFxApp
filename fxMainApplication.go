package goFxApp

import (
	"github.com/bhbosman/goConnectionManager"
	"github.com/bhbosman/goFxApp/Services/DateTimeService"
	"github.com/bhbosman/goFxApp/Services/MultiLoggerService"
	"github.com/bhbosman/goFxApp/internal"
	"github.com/bhbosman/goFxAppManager/FxServicesSlide"
	"github.com/bhbosman/goFxAppManager/service"
	UiService2 "github.com/bhbosman/goUi/UiService"
	"github.com/bhbosman/goUi/UiSlides/GoFunctionCounterSlide"
	"github.com/bhbosman/goUi/UiSlides/cmSlide"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/services/Providers"
	"github.com/cskr/pubsub"
	"go.uber.org/zap"

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
			fx.Populate(&shutDowner, &ss),
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
			goConnectionManager.ProvideCreateConnectionManager(),
			DateTimeService.ProvideDateTimeService(),
			multiLogger.ProvideMultiLogFileService(),
			service.ProvideFxManager(),
			multiLogger.InvokeMultiLogFileService(),
			internal.InvokeApps(),
			goConnectionManager.InvokeConnectionManager(),
			invokeTerminalApplicationOptions,
			service.InvokeFxManager(),
			fx.Options(option...),
			internal.InvokeApplicationContext(),
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
