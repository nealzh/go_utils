package config_utils

import (
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func InitConfigFile(configFilePath string) (*ini.File, error) {
	file, err := ini.Load(configFilePath)
	if err != nil {
		log.Fatal("配置文件读取错误，请检查文件路径:", configFilePath, err)
	}
	return file, err
}

func LoadConfigStringValues(file *ini.File, configSection string) *[]string {

	var configValue []string

	for _, key := range file.Section(configSection).Keys() {

		configValue = append(configValue, key.Value())
	}

	return &configValue
}

func LoadConfigUintValues(file *ini.File, configSection string, defaultValue uint64) *[]uint64 {

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

func LoadStringConfig(file *ini.File, configSection string, configKey string, defaultValue string) string {
	return file.Section(configSection).Key(configKey).MustString(defaultValue)
}

func LoadBoolConfig(file *ini.File, configSection string, configKey string, defaultValue bool) bool {
	return file.Section(configSection).Key(configKey).MustBool(defaultValue)
}

func LoadFloat64Config(file *ini.File, configSection string, configKey string, defaultValue float64) float64 {
	return file.Section(configSection).Key(configKey).MustFloat64(defaultValue)
}

func LoadIntConfig(file *ini.File, configSection string, configKey string, defaultValue int) int {
	return file.Section(configSection).Key(configKey).MustInt(defaultValue)
}

func LoadInt64Config(file *ini.File, configSection string, configKey string, defaultValue int64) int64 {
	return file.Section(configSection).Key(configKey).MustInt64(defaultValue)
}

func LoadUintConfig(file *ini.File, configSection string, configKey string, defaultValue uint) uint {
	return file.Section(configSection).Key(configKey).MustUint(defaultValue)
}

func LoadUint64lConfig(file *ini.File, configSection string, configKey string, defaultValue uint64) uint64 {
	return file.Section(configSection).Key(configKey).MustUint64(defaultValue)
}

func LoadDurationConfig(file *ini.File, configSection string, configKey string, defaultValue time.Duration) time.Duration {
	return file.Section(configSection).Key(configKey).MustDuration(defaultValue)
}

func LoadTimeConfig(file *ini.File, configSection string, configKey string, defaultValue time.Time) time.Time {
	return file.Section(configSection).Key(configKey).MustTime(defaultValue)
}

func LoadFileContent(filePath string) []byte {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return fileContent
}
