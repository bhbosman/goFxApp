package internal

import (
	"context"
	"go.uber.org/fx"
)

func ProvideApplicationContext() fx.Option {
	return fx.Provide(
		fx.Annotated{
			Name: "Application",
			Target: func(
				params struct {
					fx.In
				},
			) (context.Context, context.CancelFunc) {
				return context.WithCancel(context.Background())
			},
		},
	)
}
