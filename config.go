package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

type Config struct {
	VerboseMode            bool  `json:"verboseMode"`
	KeepDestinationHeaders bool  `json:"keepDestinationHeaders"`
	KeepHeader             bool  `json:"keepHeader"`
	Port                   int64 `json:"port"`
	HandleNonProxyRequests bool  `json:"handleNonproxyRequests"`
}

func ParseConfig() *Config {
	buf, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal("silverplate: cannot read ./config/config.json. do you have permission or is it missing?")
	}
	var config Config
	err = json.Unmarshal(buf, &config)
	if err != nil {
		log.Fatal("silverplate: can't parse config.json")
	}
	return &config
}

func (c *Config) PortToString() string {
	return strconv.FormatInt(c.Port, 10)
}

func (c *Config) VerboseModeToString() string {
	return strconv.FormatBool(c.VerboseMode)
}
