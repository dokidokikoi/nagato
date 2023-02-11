package config

type FileSystemConf struct {
	StoreDir string
	TempDir  string
}

const (
	fileSystemKey string = "filesystem"
)

var FlieSystemConfig = &FileSystemConf{
	StoreDir: "/tmp/objects/", TempDir: "/tmp/temp/",
}

func init() {
	fileSystemConfig := GetSpecConfig(fileSystemKey)
	if fileSystemConfig != nil {
		fileSystemConfig.Unmarshal(fileSystemConfig)
	}
}
