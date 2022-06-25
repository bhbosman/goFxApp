package Serivce

import (
	"github.com/bhbosman/goFxApp/FxAppManagerService/Serivce/internal"
	"github.com/bhbosman/gocommon/Services/IFxService"
)

type IFxManagerService interface {
	internal.IFxManager
	IFxService.IFxServices
}
