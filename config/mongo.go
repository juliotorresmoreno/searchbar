package config

import (
	"io/ioutil"
	"log"

	"encoding/json"

	"github.com/tidwall/gjson"
)

//MongoConfig Defines the structure that stores the default configuration of the mongodb database
type MongoConfig struct {
	ConnectionStr string `json:"connection_str"`
	Collection    string `json:"collection"`
}

//GetMongoConfig Gets the configuration of the current mongo database
func GetMongoConfig() MongoConfig {
	var data []byte
	response := MongoConfig{}
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatal("The mongodb database configuration file was not found")
	}
	result := gjson.Get(string(data), "mongo")
	json.Unmarshal([]byte(result), response)
	println(result.String())
	return response
}
