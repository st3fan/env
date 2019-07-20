package env

import (
	"log"
	"os"
	"strconv"
)

func MustGetString(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Fatalf("Failed to get <%s>: env var is empty", name)
	}
	return v
}

func GetString(name string, defoult string) string {
	v := os.Getenv(name)
	if v != "" {
		return v
	}
	return defoult
}

func MustGetBool(name string) bool {
	v := os.Getenv(name)
	if v == "" {
		log.Fatalf("Failed to get <%s>: env var is empty", name)
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Fatalf("Failed to parse <%s>: <%s> is not a valid boolean value", name, v)
	}
	return b
}

func GetBool(name string, defoult bool) bool {
	v := os.Getenv(name)
	if v != "" {
		b, _ := strconv.ParseBool(v)
		return b
	}
	return defoult
}

func MustGetInt(name string) int {
	v := os.Getenv(name)
	if v == "" {
		log.Fatalf("Failed to get <%s>: env var is empty", name)
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("Failed to parse <%s>: <%s> is not a valid int value", name, v)
	}
	return i

}

func GetInt(name string, defoult int) int {
	v := os.Getenv(name)
	if v != "" {
		i, _ := strconv.Atoi(v)
		return i
	}
	return defoult
}
