package conf

import "time"

type Config struct {
	App           App           `mapstructure:"app" yaml:"app"`
	Database      Database      `mapstructure:"database" yaml:"database"`
	Redis         Redis         `mapstructure:"redis" yaml:"redis"`
	Elasticsearch Elasticsearch `mapstructure:"elasticsearch" yaml:"elasticsearch"`
	Kafka         Kafka         `mapstructure:"kafka" yaml:"kafka"`
	Zap           Zap           `mapstructure:"zap" yaml:"zap"`
}

type App struct {
	Domain          string `mapstructure:"domain" yaml:"domain"`
	RunMode         string `mapstructure:"run_mode"  yaml:"run_mode"`
	RuntimeRootPath string `mapstructure:"runtime-root-path" yaml:"runtime-root-path"`
	LogSavePath     string `mapstructure:"log-save-path" yaml:"log-save-path"`
	LogSaveName     string `mapstructure:"log-save-name" yaml:"log-save-name"`
	LogFileExt      string `mapstructure:"log-file-ext" yaml:"log-file-ext"`
	TimeFormat      string `mapstructure:"time-format" yaml:"time-format"`
}

type Database struct {
	Type        string `mapstructure:"type" yaml:"type"`
	User        string `mapstructure:"user" yaml:"user"`
	Password    string `mapstructure:"password" yaml:"password"`
	Host        string `mapstructure:"host" yaml:"host"`
	Name        string `mapstructure:"name" yaml:"name"`
	TablePrefix string `mapstructure:"table-prefix" yaml:"table-prefix"`
}

type Redis struct {
	Host        string        `mapstructure:"host" yaml:"host"`
	Password    string        `mapstructure:"password" yaml:"password"`
	IdleTimeout time.Duration `mapstructure:"idle-timeout" yaml:"idle-timeout"`
}

type Elasticsearch struct {
	Hosts    []string `mapstructure:"hosts" yaml:"hosts"`
	Username string   `mapstructure:"username" yaml:"username"`
	Password string   `mapstructure:"password" yaml:"password"`
}

type Zap struct {
	LogFilePath     string `mapstructure:"log-filepath" yaml:"log-filepath"`
	LogInfoFileName string `mapstructure:"log-info-filename" yaml:"log-info-filename"`
	LogWarnFileName string `mapstructure:"log-warn-filename" yaml:"log-warn-filename"`
	LogFileExt      string `mapstructure:"log-fiile-ext" yaml:"log-fiile-ext"`
}

type Kafka struct {
	Hosts []string `mapstructure:"hosts" yaml:"hosts"`
}
