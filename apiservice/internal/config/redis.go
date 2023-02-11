package config

type RedisConf struct {
	Host     string
	Port     int
	Password string
	DB       int
}

const (
	redisKey string = "redis"
)

var RedisConfig = &RedisConf{
	Port: 6379, Host: "127.0.0.1", DB: 0, Password: "",
}

func init() {
	redisConfig := GetSpecConfig(redisKey)
	if redisConfig != nil {
		redisConfig.Unmarshal(RedisConfig)
	}
}
