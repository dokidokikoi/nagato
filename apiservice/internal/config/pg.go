package config

type PGConf struct {
	Host     string
	Port     int
	Database string `validate:"required"`
	Username string `validate:"required"`
	Password string
}

const (
	pgKey string = "postgresql"
)

var PgConfig = &PGConf{
	Port: 5432, Host: "127.0.0.1", Database: "postgres", Password: "postgres",
}

func init() {
	pgConfig := GetSpecConfig(pgKey)
	if pgConfig != nil {
		pgConfig.Unmarshal(PgConfig)
	}
}
