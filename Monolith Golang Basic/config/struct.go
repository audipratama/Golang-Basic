package config

type Config struct {
	DB   *DBConfig
	Main *MainConfig
}

type DBConfig struct {
	DB map[string]*DBSetting
}

type DBSetting struct {
	Host			string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime int
}

type MainConfig struct {
	Server MainServerConfig
}

type MainServerConfig struct {
	Environment string
	Port        int
}
