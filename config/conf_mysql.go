package config

type Mysql struct {
	Host     string `yaml:"host"`
	port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` //日志等级
}
