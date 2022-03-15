package config_utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//从配置文件中载入json字符串
func LoadConfig(path string) map[string]interface{} {

	buf, err := ioutil.ReadFile(path)

	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}

	allConfigs := make(map[string]interface{}, 0)
	err = json.Unmarshal(buf, &allConfigs)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	return allConfigs
}
