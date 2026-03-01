package src

type Config struct {
	LogoC  string
	KeyC   string
	ValueC string
	Reset  string
	Bold   string
	Sep    string
}

func SetOptions() *Config {
	cfg := &Config{
		LogoC:  "\033[34m",
		KeyC:   "\033[34m",
		ValueC: "\033[0m",
		Reset:  "\033[0m",
		Bold:   "\033[1m",
		Sep:    ":",
	}

	return cfg
}
