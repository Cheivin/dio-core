package core

import (
	"os"
)

func EnvDefault(key string, def string) (val string) {
	val = os.Getenv(key)
	if val == "" {
		val = def
	}
	return
}
