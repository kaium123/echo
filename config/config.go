package config

type DbConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	Schema string
}

type AppConfig struct {
	Name string
	Port string
}

type Config struct {
	Db  *DbConfig
	App *AppConfig
}

var config Config

func Db() *DbConfig {
	return config.Db
}

func App() *AppConfig {
	return config.App
}

func LoadConfig() {
	SetDefaultConfig()
}

func SetDefaultConfig() {
	config.Db = &DbConfig{
		Host:   "localhost",
		Port:   "3306",
		User:   "root",
		Pass:   "",
		Schema: "company",
	}

	config.App = &AppConfig{
		Name: "auth",
		Port: "8000",
	}
}
