package util

import (
	"log"
	"os"
	"strconv"
)

func MustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Panicln("ENV missing, key: " + key)
	}
	return v
}

func MustGetBool(key string) bool {
	value := os.Getenv(key)
	if value == "" {
		log.Panicln("ENV missing, key: " + key)
	}
	res, err := strconv.ParseBool(value)
	if err != nil {
		log.Panicln("ENV err: [" + key + "]\n" + err.Error())
	}
	return res
}
