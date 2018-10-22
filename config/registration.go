package config

import (
	"fmt"
	"regexp"

	"maunium.net/go/mautrix-appservice"
)

// NewRegistration returns an appservice.Registration object with appservice
// information from a configuration file
func (config *Config) NewRegistration() (*appservice.Registration, error) {
	registration := appservice.CreateRegistration()

	err := config.copyToRegistration(registration)
	if err != nil {
		return nil, err
	}

	config.AppService.ASToken = registration.AppToken
	config.AppService.HSToken = registration.ServerToken
	return registration, nil
}

// TODO: How are this and the above function different?
// NewRegistration returns an appservice.Registration object with appservice
// information from a configuration file
func (config *Config) GetRegistration() (*appservice.Registration, error) {
	registration := appservice.CreateRegistration()

	err := config.copyToRegistration(registration)
	if err != nil {
		return nil, err
	}

	registration.AppToken = config.AppService.ASToken
	registration.ServerToken = config.AppService.HSToken
	return registration, nil
}

// copyToRegistration copies information from a Config type to an
// appservice.Registration type
func (config *Config) copyToRegistration(registration *appservice.Registration) error {
	registration.ID = config.AppService.ID
	registration.URL = config.AppService.Address
	registration.RateLimited = false
	registration.SenderLocalpart = config.AppService.Bot.Username

	// Compile a regexp for the appservice's userID
	userIDRegex, err := regexp.Compile(fmt.Sprintf("^@%s:%s$",
		config.Bridge.FormatUsername("[0-9]+"),
		config.Homeserver.Domain))
	if err != nil {
		return err
	}
	registration.Namespaces.RegisterUserIDs(userIDRegex, true)
	return nil
}
