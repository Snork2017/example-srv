package handlers

type Config struct {
	Port int `env:"PORT, required" default:"8080"`
}
