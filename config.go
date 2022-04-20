package blizzapi

import (
	"fmt"
	"strings"
)

type Config struct {
	BaseUrl      string
	ClientID     string
	ClientSecret string
	Locale       string
	Region       string
	UserAgent    string
}

var staticUrlBases = map[string]string{
	"us": "us.api.blizzard.com",
	"eu": "eu.api.blizzard.com",
	"kr": "kr.api.blizzard.com",
	"tw": "tw.api.blizzard.com",
	"cn": "gateway.battlenet.com.cn",
}

// Init initializes the config struct and validates
// all options. It sets sane defaults as needed.
func (c *Config) Init() error {
	if !c.populated(c.ClientID) {
		return fmt.Errorf("a ClientID must be provided to blizzapi.Config{}")
	}

	if !c.populated(c.ClientSecret) {
		return fmt.Errorf("a ClientSecret must be provided to blizzapi.Config{}")
	}

	if !c.populated(c.Locale) {
		c.Locale = "en_US"
	}

	if !c.populated(c.Region) {
		return fmt.Errorf("a Region must be provided to blizzapi.Config{}")
	}

	c.Region = strings.ToLower(c.Region)
	c.BaseUrl = staticUrlBases[c.Region]

	if !c.populated(c.BaseUrl) {
		return fmt.Errorf("a valid region must be provided. valid regions are: us,eu,kr,tw,cn")
	}

	if !c.populated(c.UserAgent) {
		c.UserAgent = "unknown golang application"
	}
	return nil
}

// populated determines if the specified option is populated.
// it returns a bool true if the option is populated
// and bool empty if the option is not populated
func (c *Config) populated(option string) bool {
	if option == "" {
		return false
	}
	return true
}
