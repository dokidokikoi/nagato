package model

type ResumableToken struct {
	Name    string
	Size    uint
	Hash    string
	Servers []string
	Uuids   []string
}
