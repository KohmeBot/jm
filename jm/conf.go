package jm

type Config struct {
	// 请求地址
	Address string `yaml:"address" jsonschema:"description=请求地址"`
	// 冷却时间 秒
	CD int `yaml:"cd" jsonschema:"description=冷却时间 秒"`
}
