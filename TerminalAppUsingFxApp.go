package goFxApp

import (
	"context"
	"github.com/bhbosman/gocommon/uiCommon"
	"github.com/cskr/pubsub"
	"github.com/rivo/tview"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type TerminalAppUsingFxApp struct {
	FxApp       *fx.App
	TerminalApp *tview.Application
	Shutdown    fx.Shutdowner
	logger      *zap.Logger
	pubSub      *pubsub.PubSub
}

func (self *TerminalAppUsingFxApp) RunTerminalApp() {

	if err := self.FxApp.Start(context.Background()); err != nil {
		self.logger.Error("On FxApp.Start error", zap.Error(err))
		return
	}

	self.pubSub.Pub(uiCommon.NewUiStarted(true), uiCommon.UIState)
	if err := self.TerminalApp.Run(); err != nil {
		self.logger.Error("On terminal application Run error", zap.Error(err))
	}
	self.pubSub.Pub(uiCommon.NewUiStarted(false), uiCommon.UIState)

	if err := self.TerminalApp.Close(); err != nil {
		self.logger.Error("On terminal application Close error", zap.Error(err))
	}

	if err := self.FxApp.Stop(context.Background()); err != nil {
		self.logger.Error("On FxApp.Stop error", zap.Error(err))
	}
}

func NewTerminalAppUsingFxApp(
	terminalApp *tview.Application,
	shutdown fx.Shutdowner,
	fxApp *fx.App,
	logger *zap.Logger,
	pubSub *pubsub.PubSub,
) *TerminalAppUsingFxApp {
	return &TerminalAppUsingFxApp{
		FxApp:       fxApp,
		TerminalApp: terminalApp,
		Shutdown:    shutdown,
		logger:      logger,
		pubSub:      pubSub,
	}
}
