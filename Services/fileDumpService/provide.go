package fileDumpService

import (
	"fmt"
	"github.com/bhbosman/gocommon/Services/interfaces"
	"go.uber.org/fx"
	"net/url"
	"os"
)

func Provide() fx.Option {
	return fx.Provide(
		//sdfsdfds
		fx.Annotated{
			Target: func(
				params struct {
					fx.In
					ApplicationName        string `name:"ApplicationName" optional:"true"`
					UniqueReferenceService interfaces.IUniqueReferenceService
				},
			) (IFileDumpService, error) {
				dir, err := os.UserHomeDir()
				if err != nil {
					return nil, err
				}
				if params.ApplicationName == "" {
					params.ApplicationName = "UnknownApplication"
				}
				folder := fmt.Sprintf("file:///%v/ApplicationLogFiles/%v/FileDump", dir, params.ApplicationName)
				urlPath, err := url.Parse(folder)
				if err != nil {
					return nil, err
				}
				folder = urlPath.Path[1:]
				err = os.MkdirAll(folder, os.ModePerm)
				if err != nil {
					return nil, err
				}
				return newFileDumpService(params.UniqueReferenceService, folder), nil
			},
		},
	)
}
