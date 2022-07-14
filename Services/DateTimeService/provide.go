package DateTimeService

import (
	"github.com/bhbosman/gocommon/Services/IDateTimeService"
	"go.uber.org/fx"
)

func ProvideDateTimeService() fx.Option {
	return fx.Provide(
		func() (*DateTimeService, IDateTimeService.IDateTimeService) {
			v := NewDateTimeService()
			return v, v
		})
}
