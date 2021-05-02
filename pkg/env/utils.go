package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func strToInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func GetString(key string) (res string, err error) {
	res = os.Getenv(key)
	if res == "" {
		return "", fmt.Errorf("env variable with name %s not found", key)
	}
	return
}

func GetInt64(key string) (res int64, err error) {
	resString, err := GetString(key)
	if err != nil {
		return
	}
	return strToInt(resString)
}

func GetSecondsDuration(key string) (res time.Duration, err error) {
	resInt, err := GetInt64(key)
	if err != nil {
		return
	}
	return time.Duration(resInt) * time.Second, nil
}
