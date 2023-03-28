package DateTimeService

import (
	"github.com/bhbosman/gocommon/services/IDateTimeService"
	"go.uber.org/fx"
)

func ProvideDateTimeService() fx.Option {
	return fx.Provide(
		func() (*DateTimeService, IDateTimeService.IDateTimeService) {
			v := NewDateTimeService()
			return v, v
		},
	)
}
