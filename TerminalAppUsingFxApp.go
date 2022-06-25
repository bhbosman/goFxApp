package goFxApp

import (
	"context"
	"github.com/rivo/tview"
	"go.uber.org/fx"
	"os"
)

type TerminalAppUsingFxApp struct {
	FxApp       *fx.App
	TerminalApp *tview.Application
	Shutdown    fx.Shutdowner
}

func (self *TerminalAppUsingFxApp) RunTerminalApp() {
	err := self.FxApp.Start(context.Background())
	if err != nil {
		os.Exit(1)
	}
	if err = self.TerminalApp.Run(); err != nil {
		panic(err)
	}
	err = self.FxApp.Stop(context.Background())
}

func NewTerminalAppUsingFxApp(
	terminalApp *tview.Application,
	shutdown fx.Shutdowner,
	fxApp *fx.App,
) *TerminalAppUsingFxApp {
	return &TerminalAppUsingFxApp{
		FxApp:       fxApp,
		TerminalApp: terminalApp,
		Shutdown:    shutdown,
	}
}
