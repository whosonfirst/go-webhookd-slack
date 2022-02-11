// package writer implements the io.Writer interface for posting messages to Slack.
package writer

import (
	"context"
	"fmt"
	"github.com/sfomuseum/go-slack"
	"io"
)

// type SlackWriter implements the io.Writer interface for posting messages to Slack.
type SlackWriter struct {
	io.Writer
	webhook *slack.Webhook
	channel string
}

// NewSlackWriter returns a new io.Writer instance for posting messages to the Slack channel 'channel'.
func NewSlackWriter(webhook_uri string, channel string) (io.Writer, error) {

	ctx := context.Background()
	wh, err := slack.NewWebhook(ctx, webhook_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create webhook, %w", err)
	}

	wr := &SlackWriter{
		webhook: wh,
		channel: channel,
	}

	return wr, nil
}

// Write will post 'b' to the Slack channel specified when instantiating 'wr'.
func (wr *SlackWriter) Write(b []byte) (int, error) {

	m := &slack.Message{
		Channel: wr.channel,
		Text:    string(b),
	}

	ctx := context.Background()
	err := wr.webhook.Post(ctx, m)

	if err != nil {
		return 0, fmt.Errorf("Failed to post message, %w", err)
	}

	return len(b), nil
}
