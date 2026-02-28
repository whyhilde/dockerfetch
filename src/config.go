package src

type Config struct {
	LogoCol  string
	KeyCol   string
	ValueCol string
	Reset    string
	Bold     string
	Sep      string
}

func SetOptions() *Config {
	cfg := &Config{
		LogoCol:  "\033[34m",
		KeyCol:   "\033[34m",
		ValueCol: "\033[0m",
		Reset:    "\033[0m",
		Bold:     "\033[1m",
		Sep:      ":",
	}

	return cfg
}
