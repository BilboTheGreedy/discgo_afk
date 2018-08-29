package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type config struct {
	Token                  string   `json:"login_token"`
	Prefix                 string   `json:"prefix"`
	Roles                  []string `json:"allowed_roles"`
	AdminRoles             []string `json:"admin_roles"`
	DaysBeforeNotification int      `json:"days_before_notification"`
	Color                  string   `json:"color"`
}

var (
	Token                  string
	Roles                  []string
	BotID                  string
	Prefix                 string
	AdminRoles             []string
	DaysBeforeNotification int
	Color                  int
)

var cfg config

func ReadConfiguration() error {
	// Open our jsonFile
	jsonFile, err := os.Open("discord.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened discord.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &cfg)

	// we set our discord variables here

	Token = cfg.Token
	Prefix = cfg.Prefix
	Roles = cfg.Roles
	AdminRoles = cfg.AdminRoles
	DaysBeforeNotification = cfg.DaysBeforeNotification

	c, _ := strconv.ParseInt(strings.TrimPrefix(cfg.Color, "#"), 16, 64)
	Color = int(c)

	return nil
}

func SetBotID(botID string) {
	BotID = botID
	fmt.Println(botID)
}
