# go-webhookd-slack

go-webhookd support for Slack.

## go-webhookd

Before you begin please [read the go-webhookd documentation](https://github.com/whosonfirst/go-webhookd/blob/master/README.md) for an overview of concepts and principles.

## Usage

```
import (
	_ "github.com/go-webhookd-slack"
)
```

## Receivers

### Slack

The `Slack` receiver handles Webhooks sent from [Slack](https://api.slack.com/outgoing-webhooks). It does not process the message at all. It is defined as a URI string in the form of:

```
slack://
```

_This receiver has not been fully tested yet so proceed with caution._

## Transformations

### SlackText

The `SlackText` transformation will extract and return [the `text` property](https://api.slack.com/outgoing-webhooks) from a Webhook sent by Slack. It is defined as URI string in the form of:

```
slacktext://
```

## Dispatchers

### Slack

The `Slack` dispatcher will send messages to a Slack channel. It is defined as a URI string in the form of:

```
slack://?webhook={SLACK_WEBHOOK_URI}&channel={SLACK_CHANNEL}
```

#### Parameters

| Name | Value | Description | Required |
| --- | --- | --- | --- |
| webhook | string | A valid Slack Webhook URL. | yes |
| channel | string | A valid Slack channel name | yes |


Earlier versions of this package used the [whosonfirst/slackcat](https://github.com/whosonfirst/slackcat) package to send Slack messages which necessitated that `slack://` dispatcher URLs be in the form of:

```
slack://{PATH_TO_SLACKCAT.CONFG}
```

In order to preserve backwards compatibility this URL form will continue to honoured if a `slack://` dispatcher URL does not contain a `webhook` or `channel` query parameter.

#### Properties

| Name | Value | Description | Required |
| --- | --- | --- | --- |
| slackcat.conf | string | The path to a valid [slackcat](https://github.com/whosonfirst/slackcat#configuring) config file. | yes |

_Eventually you will be able to specify a plain-vanilla Slack Webhook URL but not today._

## See also

* https://github.com/whosonfirst/go-webhookd
* https://github.com/sfomuseum/go-slack