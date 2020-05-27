package slack

// https://api.slack.com/outgoing-webhooks

import (
	"bufio"
	"bytes"
	"context"
	"github.com/whosonfirst/go-webhookd/v3"
	"github.com/whosonfirst/go-webhookd/v3/transformation"
	_ "log"
	"strings"
)

func init() {

	ctx := context.Background()
	err := transformation.RegisterTransformation(ctx, "slacktext", NewSlackTextTransformation)

	if err != nil {
		panic(err)
	}
}

type SlackTextTransformation struct {
	webhookd.WebhookTransformation
	key string
}

func NewSlackTextTransformation(ctx context.Context, uri string) (webhookd.WebhookTransformation, error) {

	p := SlackTextTransformation{
		key: "text",
	}

	return &p, nil
}

func (p *SlackTextTransformation) Transform(ctx context.Context, body []byte) ([]byte, *webhookd.WebhookError) {

	buf := bytes.NewBuffer(body)
	scanner := bufio.NewScanner(buf)

	text := ""

	for scanner.Scan() {

		ln := scanner.Text()
		pair := strings.Split(ln, "=")

		if len(pair) != 2 {
			continue
		}

		if pair[0] == p.key {
			text = pair[1]
			break
		}
	}

	if text == "" {

		code := 999
		message := "Unable to parse Slack text"

		err := &webhookd.WebhookError{Code: code, Message: message}
		return nil, err
	}

	return []byte(text), nil
}
