package config

type PostgresDataBase struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
}

type JwtConfig struct {
	SecretKey           string
	ExpiresDurationHour int
}

type FileSetting struct {
	MaximunFileSizeMB int
}

type Config struct {
	Database string
	Pg       PostgresDataBase
	Jwt      JwtConfig
	Salt     string
	FileSetting
	TimeoutSecond int
}
