package errors

import (
	"errors"
	"nagato/common/code"

	myErrors "github.com/dokidokikoi/go-common/errors"
)

var (
	ErrFileNotFound = errors.New("文件未找到")
)

var (
	ApiErrFileNotFound = myErrors.ClientFailed(ErrFileNotFound.Error(), code.ErrFileNotFound)
)
