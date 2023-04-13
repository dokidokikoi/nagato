package errors

import (
	"errors"
	"nagato/common/code"

	myErrors "github.com/dokidokikoi/go-common/errors"
)

var (
	ErrFileNotFound     = errors.New("文件未找到")
	ErrFolderNotFound   = errors.New("文件夹不存在")
	ErrFolderRepeatFile = errors.New("目录下存在同名文件")
	ErrShareExpired     = errors.New("分享已过期")
	ErrShareCode        = errors.New("提取码错误")
	ErrShareNoMatters   = errors.New("未选择分享文件")
	ErrESCreateIndex    = errors.New("创建索引失败")
	ErrESCreateDoc      = errors.New("创建文档失败")
)

var (
	ApiErrFileNotFound     = myErrors.ClientFailed(ErrFileNotFound.Error(), code.ErrFileNotFound)
	ApiErrFolderNotFound   = myErrors.ClientFailed(ErrFolderNotFound.Error(), code.ErrFolderNotFound)
	ApiErrFolderRepeatFile = myErrors.ClientFailed(ErrFolderRepeatFile.Error(), code.ErrFolderRepeatFile)
	ApiErrShareExpired     = myErrors.ClientFailed(ErrShareExpired.Error(), code.ErrShareExpired)
	ApiErrShareCode        = myErrors.ClientFailed(ErrShareCode.Error(), code.ErrShareCode)
	ApiErrShareNoMatters   = myErrors.ClientFailed(ErrShareNoMatters.Error(), code.ErrShareNoMatters)
	ApiErrESCreateIndex    = myErrors.ClientFailed(ErrESCreateIndex.Error(), code.ErrESCreateIndex)
	ApiErrESCreateDoc      = myErrors.ClientFailed(ErrESCreateDoc.Error(), code.ErrESCreateDoc)
)
