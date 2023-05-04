package db

type Store interface {
	Blanks() IBlankStore
	BlankMatters() IBlankMatterStore
	Installs() IInstallStore
	Logs() ILogStore
	Matters() IMatterStore
	Shares() IShareStore
	ShareMatters() IShareMatterStore
	SmallFileCaches() ISmallFileCacheStore
	Tags() ITagStore
	Users() IUserStore
	Transaction() ITransaction
}
