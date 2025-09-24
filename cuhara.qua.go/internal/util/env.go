package util

import (
	"log"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm/logger"
)

func GetEnv(key string, defaultVal string) string {
	log.Printf("ENV: %s : %s", key, defaultVal)
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

func GetEnvAsInt(key string, defaultVal int) int {
	strVal := GetEnv(key, strconv.Itoa(defaultVal))

	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsUint32(key string, defaultVal uint32) uint32 {
	strVal := GetEnv(key, strconv.FormatUint(uint64(defaultVal), 10))

	if val, err := strconv.ParseUint(strVal, 10, 32); err == nil {
		return uint32(val)
	}

	return defaultVal
}

func GetEnvAsUint8(key string, defaultVal uint8) uint8 {
	strVal := GetEnv(key, strconv.FormatUint(uint64(defaultVal), 10))

	if val, err := strconv.ParseUint(strVal, 10, 8); err == nil {
		return uint8(val)
	}

	return defaultVal
}

func GetEnvAsBool(key string, defaultVal bool) bool {
	strVal := GetEnv(key, strconv.FormatBool(defaultVal))

	if val, err := strconv.ParseBool(strVal); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsGormLogLevel(key string, defaultVal logger.LogLevel) logger.LogLevel {
	strVal := strings.ToUpper(strings.TrimSpace(GetEnv(key, string(defaultVal))))

	switch strVal {
	case "SILENT":
		return logger.Silent
	case "ERROR":
		return logger.Error
	case "WARN":
		return logger.Warn
	case "INFO":
		return logger.Info
	default:
		return defaultVal
	}
}
