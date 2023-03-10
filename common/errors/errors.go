package errors

import (
	"errors"
	"nagato/common/code"

	myErrors "github.com/dokidokikoi/go-common/errors"
)

var (
	ErrFileNotFound   = errors.New("文件未找到")
	ErrFolderNotFound = errors.New("文件夹不存在")
)

var (
	ApiErrFileNotFound   = myErrors.ClientFailed(ErrFileNotFound.Error(), code.ErrFileNotFound)
	ApiErrFolderNotFound = myErrors.ClientFailed(ErrFolderNotFound.Error(), code.ErrFolderNotFound)
)
