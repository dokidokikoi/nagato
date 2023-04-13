package db

type Store interface {
	Blank() IBlankStore
	Resource() IResourceStore
}
