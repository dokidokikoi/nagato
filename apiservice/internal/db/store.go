package db

import "fmt"

type Store interface {
	Blanks() IBlankStore
	BlankMatters() IBlankMatterStore
	Bridges() IBridgeStore
	Installs() IInstallStore
	Logs() ILogStore
	Matters() IMatterStore
	Shares() IShareStore
	SmallFileCaches() ISmallFileCacheStore
	Tags() ITagStore
	Users() IUserStore
	Transaction() ITransaction
}

var storePointer Store

func GetStoreFactory() (Store, error) {
	if storePointer == nil {
		return nil, fmt.Errorf("数据层未初始化")
	}
	return storePointer, nil
}

func SetStoreFactory(factory Store) {
	storePointer = factory
}
