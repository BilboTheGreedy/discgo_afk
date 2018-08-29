package main

import (
	"fmt"

	config "./config"

	bot "./bot"
)

func main() {

	err := config.ReadConfiguration()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(config.Token)
	bot.Start()

	<-make(chan struct{})

	return
}
