package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}

type System struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}
