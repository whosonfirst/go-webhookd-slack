package slack

// type Message defines a struct for messages posted to Slack
type Message struct {
	// The name of the channel to post the message to.
	Channel string `json:"channel"`
	// The body of the message to post to a channel.
	Text string `json:"text"`
}
