package config

type FileSystemConfig struct {
	StoreDir string
	TempDir  string
}

const (
	fileSystemKey string = "filesystem"
)

func GetFileSystemInfo() FileSystemConfig {
	fileSystemConfig := &FileSystemConfig{
		StoreDir: "/tmp/objects/", TempDir: "/tmp/temp/",
	}
	conf := GetSpecConfig(fileSystemKey)
	if conf != nil {
		conf.Unmarshal(fileSystemConfig)
	}

	return *fileSystemConfig
}
