package main

var (
	// 版本信息
	appVersion bool // 控制是否显示版本
	APPVersion = "v0.0.2"
	BuildTime  = "2006-01-02 15:04:05"
	GoVersion  = "go1.21"
	GitCommit  = "xxxxxxxxxxx"
	ConfigFile = "config.yaml"
	config     *Config

	AesKey  = "wzFdVviHTKraaPRWEa9bFLLzTkddtUNY"
	AesData string // 用于存储明文
)

type Service struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

type Config struct {
	Service Service `json:"service" yaml:"service"`
	Encrypt bool    `yaml:"encrypt"`
	Token   string  `yaml:"token"`
	Cron    string  `yaml:"cron"`
}
