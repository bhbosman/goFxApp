package internal

import (
	"context"
	"go.uber.org/fx"
)

func InvokeApplicationContext() fx.Option {
	return fx.Invoke(
		func(
			params struct {
				fx.In
				CancelFunc context.CancelFunc `name:"Application"`
				Lifecycle  fx.Lifecycle
			},
		) error {
			params.Lifecycle.Append(
				fx.Hook{
					OnStart: nil,
					OnStop: func(ctx context.Context) error {
						params.CancelFunc()
						return nil
					},
				},
			)
			return nil
		},
	)
}
