package internal

import (
	"context"
	"github.com/bhbosman/gocommon/messages"
	"go.uber.org/fx"
)

func InvokeApps() fx.Option {
	return fx.Invoke(
		func(params struct {
			fx.In
			Lifecycle fx.Lifecycle
			Apps      []*fx.App                    `group:"Apps"`
			FnApps    []messages.CreateAppCallback `group:"Apps"`
		}) error {
			for _, item := range params.Apps {
				localApp := item
				params.Lifecycle.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return localApp.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return localApp.Stop(ctx)
					},
				})
			}
			return nil
		},
	)
}
