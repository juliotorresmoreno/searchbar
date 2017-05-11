package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//RedisConfig Defines the structure that stores the default configuration of the mongodb database
type RedisConfig struct {
	Host string `json:"host"`
}

//GetRedisConfig Gets the configuration of the current mongo database
func GetRedisConfig() RedisConfig {
	var data []byte
	response := RedisConfig{}
	data, err := ioutil.ReadFile("/config/RISP_SearchLocalRedis.json")
	if err != nil {
		log.Fatal("The redis database configuration file was not found")
	}
	json.Unmarshal(data, &response)
	return response
}
