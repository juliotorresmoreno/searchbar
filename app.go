package main

import (
	"flag"
	"fmt"
	"os"

	"encoding/json"

	"github.com/juliotorresmoreno/searchbar-rebuild/config"
)

func main() {
	/*	.default('p', port)
		.alias('p', 'port')
		.describe('p', 'port to host search features on')
		.default('m', './config/RISP_SearchLocalMongos.json')
		.alias('m', 'mongo')
		.describe('m', 'options file defining mongo connection. defaults to ./config/RISP_SearchLocalMongos.json')
		.default('r', './config/RISP_SearchLocalRedis.json')
		.alias('r', 'redis')
		.describe('r', 'options file defining redis connection. defaults to ./config/RISP_SearchLocalRedis.json')
		.boolean('i')
		.alias('i', 'init')
		.describe('i', 'initialize redis with a fresh copy of LocationData before starting')
		.boolean('d')
		.alias('d', 'drop')
		.describe('d', 'drop the existing data-set prior to loading')
		.default('s', 'IMLS')
		.alias('s', 'source')
		.describe('s', 'MLS resource used for prefixing location data')
		.alias('u', 'Users')
		.describe('u', 'Sync User collection')
		.default('j', false)
		.alias('j', 'json')
		.describe('j', 'import locations from a JSON file')*/
	port := os.Getenv("PORT")
	if port == "" {
		port = "8099"
	}
	_port := flag.String("p", port, "a string")

	mongoConfig := config.GetMongoConfig()
	_mongoConfig := flag.String("m", "", "a string")
	if *_mongoConfig == "" {
		json.Unmarshal([]byte(*_mongoConfig), &mongoConfig)
	}

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
	fmt.Println("port:", *_port)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
