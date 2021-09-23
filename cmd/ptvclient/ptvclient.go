package main

import (
	"errors"
	"flag"
	"github.com/alecthomas/gometalinter/_linters/src/gopkg.in/yaml.v2"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"github.com/surdaft/pterodactyl-valheim-discord-egg"
	"github.com/surdaft/pterodactyl-valheim-discord-egg/pterodactyl"
	"github.com/surdaft/pterodactyl-valheim-discord-egg/pterodactyl/responses"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "configPath", "", "Path to config file")
	flag.Parse()

	if configPath == "" {
		log.Fatal(errors.New("Config path missing"))
	}

	configFileData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panic(err)
	}

	var config pterodactyl_valheim_discord_egg.Config

	err = yaml.Unmarshal(configFileData, &config)
	if err != nil {
		log.Panic(err)
	}

	var client *pterodactyl.Client
	client = &pterodactyl.Client{
		Config: config.Pterodactyl,
	}

	clientResponse := &responses.ClientResponse{}
	client.Get("/api/client", &clientResponse)

	for _, i2 := range clientResponse.Data {
		log.Printf("%v", i2.Attributes.Name)
	}
}
