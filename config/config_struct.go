package config

type DbConfig struct {
	Dsn string
}

type Configuration struct {
	Server ServerConfiguration
	Auth   AuthConfig
	Db     DbConfig
}

type ServerConfiguration struct {
	Port  int
	Names []string
}

type AuthConfig struct {
	ClientId string
	JwksUrl  string
	Audience string
}
