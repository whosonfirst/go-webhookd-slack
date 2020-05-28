# go-webhookd-slack

## Important

Work in progress.

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
slack://{PATH_TO_SLACKCAT.CONFG}
```

#### Properties

| Name | Value | Description | Required |
| --- | --- | --- | --- |
| slackcat.conf | string | The path to a valid [slackcat](https://github.com/whosonfirst/slackcat#configuring) config file. | yes |

_Eventually you will be able to specify a plain-vanilla Slack Webhook URL but not today._

## See also

* https://github.com/whosonfirst/go-webhookd