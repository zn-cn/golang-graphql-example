package conf

// baseConfig
type baseConfig struct {
	Env      string `toml:"env"`
	LogLevel string `toml:"log_level"`
	APP      app    `toml:"app"`
	DB       db     `toml:"db"`
}

type app struct {
	Name    string `toml:"name"`
	Address string `toml:"address"`
}

type db struct {
	DBName   string `toml:"db_name"`
	UserName string `toml:"user_name"`
	PWD      string `toml:"pwd"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
}
