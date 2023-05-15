package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"regexp"
	"strings"
)

const envVariableRegexPattern = `\$\{(.+?)(?:\|\|(.*?))?\}` // ${SERVER_PORT||8080}

func InitAndReadConfig() Configuration {

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	replaceEnvVariables()

	var configuration Configuration
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return configuration
}

func replaceEnvVariables() {
	for _, key := range viper.AllKeys() {
		value := viper.Get(key)

		switch t := value.(type) {
		case string:
			viper.Set(key, replaceEnvVariable(t))
		case []any:
			viper.Set(key, replaceEnvVariablesForList(t))
		}

		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		err := viper.BindEnv(key, envKey)
		if err != nil {
			log.Fatal("config: unable to bind env: " + err.Error())
		}
	}
}

func replaceEnvVariable(value string) string {
	re := regexp.MustCompile(envVariableRegexPattern)
	parts := re.FindStringSubmatch(value)

	if parts == nil {
		return value
	}
	envValue := os.Getenv(parts[1])
	if len(envValue) > 0 {
		return envValue
	} else {
		if parts[2] != "" {
			return parts[2]
		} else {
			log.Fatalf("Unable to resolve env variable %s and there isn't default value", parts[1])
		}
	}
	return value
}

func replaceEnvVariablesForList(list []any) []any {
	for i, value := range list {
		switch t := value.(type) {
		case string:
			list[i] = replaceEnvVariable(t)
		case []any:
			list[i] = replaceEnvVariablesForList(t)
		}
	}
	return list
}
