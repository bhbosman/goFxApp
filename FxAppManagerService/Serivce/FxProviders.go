package Serivce

import (
	"context"
	internal2 "github.com/bhbosman/goFxApp/FxAppManagerService/Serivce/internal"
	"github.com/bhbosman/gocommon/messages"
	"github.com/cskr/pubsub"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func InvokeFxManager() fx.Option {
	return fx.Options(
		fx.Invoke(
			func(
				params struct {
					fx.In
					Lifecycle        fx.Lifecycle
					FxManagerService IFxManagerService
				},
			) error {
				hook := fx.Hook{
					OnStart: params.FxManagerService.OnStart,
					OnStop:  params.FxManagerService.OnStop,
				}
				params.Lifecycle.Append(hook)
				return nil
			},
		),
	)
}

func ProvideFxManager() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Target: func(
					params struct {
						fx.In
						OnData             internal2.OnDataCallback
						ApplicationContext context.Context `name:"Application"`
						PubSub             *pubsub.PubSub  `name:"Application"`
					},
				) (IFxManagerService, error) {
					return internal2.NewFxManagerService(
						params.ApplicationContext,
						params.OnData,
					)
				},
			},
		),
		fx.Provide(
			fx.Annotated{
				Target: func(
					params struct {
						fx.In
						PubSub *pubsub.PubSub               `name:"Application"`
						FnApps []messages.CreateAppCallback `group:"Apps"`
						Logger *zap.Logger
					},
				) internal2.OnDataCallback {
					return func(applicationContext context.Context) (internal2.IFxManagerData, error) {
						return internal2.NewData(
							applicationContext,
							params.FnApps,
							params.PubSub,
							params.Logger.Named("FxServiceData"),
						)
					}
				},
			},
		),
	)
}
