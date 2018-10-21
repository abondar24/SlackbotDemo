package util

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
)

type SlackClient struct {
	api *slack.Client
}

const SendToChannel = "babahchannel"

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
	user := readUserInfo(client.api, userEmail)
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)

}

func readUserInfo(api *slack.Client, userEmail string) *slack.User {
	user, err := api.GetUserByEmail(userEmail)
	if err != nil {
		log.Fatalln(err)
	}

	return user
}

func (client *SlackClient) ReadChannels() {
	channels := readChannels(client.api)

	for _, channel := range channels {
		fmt.Printf("ID: %s, Name: %s, Members: %d\n", channel.ID, channel.Name, channel.NumMembers)
	}

}

func readChannels(api *slack.Client) []slack.Channel {
	channels, err := api.GetChannels(false)
	if err != nil {
		log.Fatalln(err)
	}

	return channels

}

func (client *SlackClient) SendMessageToChannel() {
	attachment := slack.Attachment{
		Pretext: "Hey.",
		Text:    " How are you? I am a bot",
	}

	channels := readChannels(client.api)

	var channelId string
	for _, channel := range channels {
		if channel.Name == SendToChannel {
			channelId = channel.ID
		}
	}

	if channelId == "" {
		log.Fatalln("channel doesn't exist")
		return
	}

	channelId, timestamp, err := client.api.PostMessage(channelId, slack.MsgOptionText("Demo Message", false),
		slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Message successfully sent to channel %s at %s\n", channelId, timestamp)

}

func (client *SlackClient) SendMessageToUser(userEmail string) {
	user := readUserInfo(client.api, userEmail)

	_, _, channelId, err := client.api.OpenIMChannel(user.ID)
	if err != nil {
		log.Fatalln(err)
	}

	_, timestamp, err := client.api.PostMessage(channelId, slack.MsgOptionText("Hey man", false))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Message successfully sent to user %s at %s\n", user.ID, timestamp)
}

func (client *SlackClient) CheckBilling(userEmail string) {
	user := readUserInfo(client.api, userEmail)

	workspaceClient := slack.New("workspace token")
	billingActive, err := workspaceClient.GetBillableInfo(user.ID)
	if err != nil {
		log.Println("ss1")
		log.Fatalln(err)
	}

	fmt.Printf("User ID: %s, BillingActive: %v\n\n\n", user.ID, billingActive[user.ID])

	billingActive, err = workspaceClient.GetBillableInfoForTeam()
	if err != nil {
		log.Fatalln(err)
	}

	for id, value := range billingActive {
		fmt.Printf("User ID: %s, BillingActive: %v\n\n\n", id, value)
	}
}

func (client *SlackClient) Stars() {
	params := slack.NewStarsParameters()
	accessClient := slack.New("workspace token")
	starredItems, _, err := accessClient.GetStarred(params)
	if err != nil {
		log.Fatalln(err)
	}

	for _, s := range starredItems {
		var desc string
		switch s.Type {
		case slack.TYPE_MESSAGE:
			desc = s.Message.Text
		case slack.TYPE_FILE:
			desc = s.File.Name
		case slack.TYPE_FILE_COMMENT:
			desc = s.File.Name + " - " + s.Comment.Comment
		case slack.TYPE_CHANNEL, slack.TYPE_IM, slack.TYPE_GROUP:
			desc = s.Channel
		}

		fmt.Println("Starred ", s.Type, ":", desc)

	}

}
