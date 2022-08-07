package fileDumpService

import "io"

type IFileDumpService interface {
	CreateTempFile(s string) (io.WriteCloser, string, error)
}
