package config_utils

import (
	"gopkg.in/ini.v1"
	"log"
	"strconv"
	"time"
)

func InitIniConfigFile(configFilePath string) (*ini.File, error) {
	file, err := ini.Load(configFilePath)
	if err != nil {
		log.Fatal("配置文件读取错误，请检查文件路径:", configFilePath, err)
	}
	return file, err
}

func LoadIniConfigStringValues(file *ini.File, configSection string) *[]string {

	var configValue []string

	for _, key := range file.Section(configSection).Keys() {

		configValue = append(configValue, key.Value())
	}

	return &configValue
}

func LoadIniConfigUintValues(file *ini.File, configSection string, defaultValue uint64) *[]uint64 {

	var configValue []uint64

	for _, key := range file.Section(configSection).Keys() {

		v, err := strconv.ParseUint(key.Value(), 10, 64)

		if err != nil {
			v = defaultValue
		}

		configValue = append(configValue, v)
	}

	return &configValue
}

func LoadIniStringConfig(file *ini.File, configSection string, configKey string, defaultValue string) string {
	return file.Section(configSection).Key(configKey).MustString(defaultValue)
}

func LoadIniBoolConfig(file *ini.File, configSection string, configKey string, defaultValue bool) bool {
	return file.Section(configSection).Key(configKey).MustBool(defaultValue)
}

func LoadIniFloat64Config(file *ini.File, configSection string, configKey string, defaultValue float64) float64 {
	return file.Section(configSection).Key(configKey).MustFloat64(defaultValue)
}

func LoadIniIntConfig(file *ini.File, configSection string, configKey string, defaultValue int) int {
	return file.Section(configSection).Key(configKey).MustInt(defaultValue)
}

func LoadIniInt64Config(file *ini.File, configSection string, configKey string, defaultValue int64) int64 {
	return file.Section(configSection).Key(configKey).MustInt64(defaultValue)
}

func LoadIniUintConfig(file *ini.File, configSection string, configKey string, defaultValue uint) uint {
	return file.Section(configSection).Key(configKey).MustUint(defaultValue)
}

func LoadIniUint64lConfig(file *ini.File, configSection string, configKey string, defaultValue uint64) uint64 {
	return file.Section(configSection).Key(configKey).MustUint64(defaultValue)
}

func LoadIniDurationConfig(file *ini.File, configSection string, configKey string, defaultValue time.Duration) time.Duration {
	return file.Section(configSection).Key(configKey).MustDuration(defaultValue)
}

func LoadIniTimeConfig(file *ini.File, configSection string, configKey string, defaultValue time.Time) time.Time {
	return file.Section(configSection).Key(configKey).MustTime(defaultValue)
}
