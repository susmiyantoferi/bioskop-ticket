package config

type Configuration struct {
	App      AppConfig      `mapstructure:"app"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Jwt      JWTConfig      `mapstructure:"jwt"`
}

type AppConfig struct {
	Port        string `mapstructure:"port"`
	Environment string `mapstructure:"environment"`
}

type PostgresConfig struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	IdleConnect int    `mapstructure:"idleconnect"`
	MaxConnect  int    `mapstructure:"maxconnect"`
	LifeConnect int    `mapstructure:"lifeconnect"`
}

type LoggerConfig struct {
	Level int    `mapstructure:"level"`
	Path  string `mapstructure:"path"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}
