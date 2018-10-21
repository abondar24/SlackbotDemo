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

	client := util.InitClient(props["botUserToken"])
	client.ReadGroups()
	client.ReadUserInfo(props["adminEmail"])
	client.ReadChannels()
	client.SendMessageToChannel()
	client.SendMessageToUser(props["adminEmail"])

	client.React(props["authToken"])
	client.HandleSlashCommand(props["verToken"])
	//client.CheckBilling(props["adminEmail"])
	//client.Stars(props["botAccessToken"])
}
