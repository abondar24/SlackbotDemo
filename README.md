# SlackbotDemo
Go based slackbot

Demo shows some features of Slack bots.

### Features

- Read groups
- Read user info by email
- Read channels
- Send message to selected channel
- Send message to selected user 
- Stared items(not working as current lib version doesn't support workspa ce tokens)
- Billable info(not working as current lib version doesn't support workspa ce tokens)
- Send reactions
- Handle slash commands(requried a tunneling server e.g [ngrock](https://ngrok.com/])

### Build and run

```yaml
go get
go build

./main
```
Check token on bot admin page in Slack