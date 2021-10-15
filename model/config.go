package model

type Config struct {
	Title    string `toml:"title"`
	LogLevel string `toml:"log_level"`
	Redis    redis  `toml:"redis"`
	DB       mysql  `toml:"mysql"`
	Server   server `toml:"server"`
}

type redis struct {
	Host     string
	Port     string
	Password string
	Database int
}

type mysql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

type server struct {
	GRPC string `toml:"grpc"`
}
