package config

type Config struct {
	DatabaseURL string
}

const (
	ErrorLoadEnv = "Error loading .env file"
)

func LoadConfig() *Config {
	return &Config{
		DatabaseURL: "host=localhost port=5430 user=postgres password=sql091233 dbname=Todolist sslmode=disable",
	}
}
