package fileDumpService

import (
	"github.com/bhbosman/gocommon/Services/interfaces"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type fileDumpService struct {
	UniqueReferenceService interfaces.IUniqueReferenceService
	folder                 string
}

func (self *fileDumpService) CreateTempFile(s string) (io.WriteCloser, string, error) {
	s = strings.Replace(s, "/", "", -1)
	dd := self.UniqueReferenceService.Next(s)
	fileName := filepath.Join(self.folder, dd)
	open, err := os.Create(fileName)
	if err != nil {
		return nil, "", err
	}
	return open, fileName, nil
}

func newFileDumpService(
	uniqueReferenceService interfaces.IUniqueReferenceService,
	folder string,
) IFileDumpService {
	return &fileDumpService{
		UniqueReferenceService: uniqueReferenceService,
		folder:                 folder,
	}
}
