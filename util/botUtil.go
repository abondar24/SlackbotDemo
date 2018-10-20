package util

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
)

type SlackClient struct {
	api *slack.Client
}

func InitClient(token string) *SlackClient {
	api := slack.New(token)
	return &SlackClient{api: api}
}

func (client *SlackClient) ReadGroups() {
	groups, err := client.api.GetGroups(false)
	if err != nil {
		log.Fatalln(err)
	}

	if len(groups) == 0 {
		fmt.Println("No groups found")
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)

	}
}

func (client *SlackClient) ReadUserInfo(userEmail string) {
	user, err := client.api.GetUserByEmail(userEmail)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)

}
