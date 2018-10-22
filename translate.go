package main

import (
	"fmt"
	"log"

	"maunium.net/go/mautrix-appservice"
	"maunium.net/go/mauflag"
	"github.com/anoadragon453/matrix-appservice-translate/config"
)

const (
	configFilepath := "config.yaml"
)

// createAppService creates an appservice with the specified config info
func createAppService(config *config.Config) (*appservice.AppService, error) {
	// Create and configure an appservice instance
	as := appservice.Create()
	as.HomeserverDomain = config.Homeserver.Domain
	as.HomeserverURL = config.Homeserver.Address
	as.Host.Hostname = config.AppService.Hostname
	as.Host.Port = config.AppService.Port

	// Apply registration information to the appservice
	var err error
	as.Registration, err = config.GetRegistration()
	return as, err
}

func main() {
	// Read config file
	botConfig, err := config.Load(configFilepath)
	if err != nil {
		log.Fatal("Unable to load config file:", err)
	}

	// Create new appservice instance
	as, err := createAppService(botConfig)
	if err != nil {
		log.Fatal("Unable to create appservice:", err)
	}

	appservice.Start()
}
