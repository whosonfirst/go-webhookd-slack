package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// type Webhook defines a struct for sending messages to a Slack Webhook URL.
type Webhook struct {
	client *http.Client
	url    string
}

// NewWebhook returns a new `Webhook` instance for posting messages to 'url'.
func NewWebhook(ctx context.Context, url string) (*Webhook, error) {

	cl := &http.Client{}

	wh := &Webhook{
		client: cl,
		url:    url,
	}

	return wh, nil
}

// > curl -X POST -H 'Content-type: application/json' --data '{"text":"Hello, World!"}' https://hooks.slack.com/services/.../.../...

// Post will post 'm' to the Webhook URL associated with 'wh'.
func (wh *Webhook) Post(ctx context.Context, m *Message) error {

	select {
	case <-ctx.Done():
		return nil
	default:
		// pass
	}

	enc, err := json.Marshal(m)

	if err != nil {
		return fmt.Errorf("Failed to marshal message, %w", err)
	}

	buf := bytes.NewBuffer(enc)

	req, err := http.NewRequest("POST", wh.url, buf)

	if err != nil {
		return fmt.Errorf("Failed to create new request, %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	rsp, err := wh.client.Do(req)

	if err != nil {
		return fmt.Errorf("Failed to POST webhook, %w", err)
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code %d", rsp.StatusCode)
	}

	return nil
}
