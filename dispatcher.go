package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sfomuseum/go-slack/writer"
	"github.com/whosonfirst/go-webhookd/v3"
	"github.com/whosonfirst/go-webhookd/v3/dispatcher"
	"io"
	"net/url"
	"os"
)

func init() {

	ctx := context.Background()
	err := dispatcher.RegisterDispatcher(ctx, "slack", NewSlackDispatcher)

	if err != nil {
		panic(err)
	}
}

// For backwards compatibility

type SlackcatConfig struct {
	WebhookUrl string `json:"webhook_url"`
	Channel    string `json:"channel"`
	Username   string `json:"username"`
}

type SlackDispatcher struct {
	webhookd.WebhookDispatcher
	writer io.Writer
}

func NewSlackDispatcher(ctx context.Context, uri string) (webhookd.WebhookDispatcher, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %v", err)
	}

	q := u.Query()

	wh_uri := q.Get("webhook")
	wh_channel := q.Get("channel")

	// Handle older configs which assumed this package was using whosonfirst/slackcat
	// to send messages

	if wh_uri == "" || wh_channel == "" {

		config_path := u.Path
		config_r, err := os.Open(config_path)

		if err != nil {
			return nil, fmt.Errorf("Failed to open %s, %w", config_path, err)
		}

		defer config_r.Close()

		var cfg *SlackcatConfig

		dec := json.NewDecoder(config_r)
		err = dec.Decode(&cfg)

		if err != nil {
			return nil, fmt.Errorf("Failed to decode %s, %w", config_path, err)
		}

		wh_uri = cfg.WebhookUrl
		wh_channel = cfg.Channel
	}

	wr, err := writer.NewSlackWriter(wh_uri, wh_channel)

	if err != nil {
		return nil, err
	}

	slack := SlackDispatcher{
		writer: wr,
	}

	return &slack, nil
}

func (sl *SlackDispatcher) Dispatch(ctx context.Context, body []byte) *webhookd.WebhookError {

	select {
	case <-ctx.Done():
		return nil
	default:
		// pass
	}

	_, err := sl.writer.Write(body)

	if err != nil {
		code := 999
		message := err.Error()

		err := &webhookd.WebhookError{Code: code, Message: message}
		return err
	}

	return nil
}
