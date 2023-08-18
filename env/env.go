package env

import "os"

var (
	envs = make(map[string]string)
)

func GetEnv(key string) (value string, ok bool) {
	value, ok = envs[key]
	if !ok {
		value, ok = os.LookupEnv(key)
	}
	return
}

func GetAll() map[string]string {
	return envs
}

func GetEnvOrDefault(key string, defaultValue string) string {
	value, ok := GetEnv(key)
	if ok {
		return value
	}
	return defaultValue
}
