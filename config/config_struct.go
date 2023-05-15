package config

type Configuration struct {
	Server ServerConfiguration
	Auth   AuthConfig
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
