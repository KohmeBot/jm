package jm

type Config struct {
	// 请求地址
	Address string `yaml:"address"`
	// 冷却时间 秒
	CD int `yaml:"cd"`
}
