package shittystartup

import "flag"

type Config struct {
	StaticURL string
	Port      int
}

// The config constructor function
// It returns the config by value (non-pointer) because it is a lightweight
// struct of integers and strings
func ParseConfig() Config {
	config := Config{}

	// Allow the config values to be set from the command line
	flag.StringVar(
		&config.StaticURL,
		"staticurl",
		"/static/",
		"static file URL",
	)
	flag.IntVar(&config.Port, "port", 9001, "port number")

	// Don't forget to parse!
	flag.Parse()

	return config
}
