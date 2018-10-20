package main

import (
	"github.com/abondar24/SlackbotDemo/util"
	"log"
)

func main() {
	props, err := util.ReadProperties()
	if err != nil {
		log.Fatalln(err)
	}

	client := util.InitClient(props["token"])
	client.ReadGroups()

	client.ReadUserInfo(props["adminEmail"])
	client.ReadChannels()
	client.SendMessageToChannel()
	client.SendMessageToUser(props["adminEmail"])
}
