package config

type DB struct {
	Host     string `env:"DB_HOST, required"`     // the database host to connect to
	Port     string `env:"DB_PORT, required"`     // the database port to connect to
	User     string `env:"DB_USER, required"`     // the user to connect to the database with
	Password string `env:"DB_PASSWORD, required"` // the password to connect to the database with
	Name     string `env:"DB_NAME, required"`     // the name of the database to connect to
}
