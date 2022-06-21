package FxAppManagerService

import (
	"context"
	"github.com/bhbosman/gocommon/Services/IDataShutDown"
	"github.com/bhbosman/gocommon/Services/IFxService"
	"github.com/bhbosman/gocommon/Services/ISendMessage"
)

type IFxManager interface {
	StopAll(ctx context.Context) error
	StartAll(ctx context.Context) error
	Stop(ctx context.Context, name ...string) error
	Start(ctx context.Context, name ...string) error
}

type IFxManagerService interface {
	IFxManager
	IFxService.IFxServices
}

type IFxManagerData interface {
	IFxManager
	IDataShutDown.IDataShutDown
	ISendMessage.ISendMessage
}
