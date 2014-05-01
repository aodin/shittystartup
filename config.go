package shittystartup

import ()

type Config struct {
	StaticURL string
	Port      int64
}

// The config constructor function
// It returns the config by value (non-pointer) because it is a lightweight
// struct of integers and strings
func ParseConfig() Config {
	config := Config{
		StaticURL: "/static/",
		Port:      9001,
	}
	return config
}
