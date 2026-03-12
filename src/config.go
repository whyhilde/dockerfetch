package src

type Config struct {
	LogoC    string
	KeyC     string
	ValueC   string
	KeyWidth int
	Reset    string
	Bold     string
	Sep      string
}

func SetOptions() *Config {
	cfg := &Config{
		LogoC:    "blue",
		KeyC:     "blue",
		ValueC:   "normal",
		KeyWidth: 14,
		Sep:      ":",
	}

	return cfg
}
