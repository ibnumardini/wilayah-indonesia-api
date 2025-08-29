package env

import (
	"os"
	"strconv"
)

func Get(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func GetAsBool(name string, defaultVal bool) bool {
	valStr := Get(name, "")
	var truthy = []string{
		"TRUE",
		"True",
		"true",
		"1",
	}
	for _, val := range truthy {
		if val == valStr {
			return true
		}
	}

	return false
}

func GetAsFloat64(name string, defaultVal float64) float64 {
	valueStr := Get(name, "")

	if val, err := strconv.ParseFloat(valueStr, 32); err == nil {
		return val
	}

	return defaultVal
}

func GetAsInt(name string, defaultVal int) int {
	valueStr := Get(name, "")

	if val, err := strconv.Atoi(valueStr); err == nil {
		return val
	}

	return defaultVal
}
