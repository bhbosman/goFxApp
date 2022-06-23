package UiService

import (
	"context"
	"github.com/bhbosman/goFxApp/ui"
	"github.com/rivo/tview"
	"go.uber.org/fx"
)

func InvokeTerminalApplication() fx.Option {
	return fx.Options(
		fx.Invoke(
			func(
				params struct {
					fx.In
					Lifecycle fx.Lifecycle
					UiApp     IUiService
				},
			) {
				hook := fx.Hook{
					OnStart: params.UiApp.OnStart,
					OnStop:  params.UiApp.OnStop,
				}
				params.Lifecycle.Append(hook)
			},
		),
		fx.Invoke(
			func(
				params struct {
					fx.In
					Lifecycle       fx.Lifecycle
					PrimitiveCloser ui.IPrimitiveCloser
				},
			) {
				hook := fx.Hook{
					OnStart: nil,
					OnStop: func(ctx context.Context) error {
						return params.PrimitiveCloser.Close()
					},
				}
				params.Lifecycle.Append(hook)
			},
		),
		fx.Invoke(
			func(
				params struct {
					fx.In
					UiApp           IUiService
					App             *tview.Application
					PrimitiveCloser ui.IPrimitiveCloser
				},
			) error {
				params.App.SetRoot(params.PrimitiveCloser, true).EnableMouse(true)
				return nil
			},
		),
	)
}
