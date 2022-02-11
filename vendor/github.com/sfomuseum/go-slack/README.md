# go-slack

SFO Museum's opinionated Go package for doing things with Slack.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/sfomuseum/go-slack.svg)](https://pkg.go.dev/github.com/sfomuseum/go-slack)

## Example

```
package main

import (
	"context"
	"github.com/sfomuseum/go-slack"
)

func main() {

	ctx := context.Background()

	webhook_uri := "https://hooks.slack.com/services/.../.../..."
	channel := "test"
	text := "hello world"

	wh, _ := slack.NewWebhook(ctx, webhook_uri)

	m := &slack.Message{
		Channel: channel,
		Text:    text,
	}

	wh.Post(ctx, m)
}
```

_Error handling omitted for the sake of brevity._

## Tools

```
$> make cli
go build -mod vendor -o bin/to-slack cmd/to-slack/main.go
```

### to-slack

Post a message to a Slack channel. The principal difference between `to-slack` and other similar tools is the use of [Go Cloud runtimevar URIs](https://gocloud.dev/howto/runtimevar) to define Slack Webhook URIs. This allows these otherwise sensitive values to be stored and retrieved from a variety of different storage mechanisms.

```
> ./bin/to-slack -h
Post a message to a Slack channel.
Usage:
	 ./bin/to-slack [options] message
Valid options are:
  -channel string
    	A valid Slack channel to post to.
  -stdin
    	Read input from STDIN
  -webhook-uri string
    	A valid gocloud.dev/runtimevar URI containing a Slack webhook URL to post to.
```

For example:

```
$> ./bin/to-slack -channel test -webhook-uri 'constant://?val=https://hooks.slack.com/services/.../.../...' testing
```

Or posting data read from `STDIN`:

```
$> echo "wub wub wub" | ./bin/to-slack -channel test -webhook-uri 'constant://?val=https://hooks.slack.com/services/.../.../...' -stdin
```

Or posting data reading Webhook URL information from an AWS Parameter Store secret:

```
$> echo "wub wub wub" | ./bin/to-slack -channel test -webhook 'awsparamstore://{SECRET_NAME}?region={REGION}&credentials={CREDENTIALS}' -stdin
```

#### Runtimevar(s)

The following Go Cloud `runtimevar` services are supported by the `to-slack` tool:

* [AWS Parameter Store](https://gocloud.dev/howto/runtimevar/#awsps)
* [Local](https://gocloud.dev/howto/runtimevar/#local)

It is possible to load runtimevar data from AWS Parameter Store using [aaronland/go-aws-session](https://github.com/aaronland/go-aws-session) -style credential strings. For example:

```
awsparamstore://{KEY}?region={REGION}&credentials={CREDENTIALS}
```

Credentials for AWS sessions are defined as string labels. They are:

| Label | Description |
| --- | --- |
| `env:` | Read credentials from AWS defined environment variables. |
| `iam:` | Assume AWS IAM credentials are in effect. |
| `{AWS_PROFILE_NAME}` | Use the profile from the default AWS credentials location. |
| `{AWS_CREDENTIALS_PATH}:{AWS_PROFILE_NAME}` | Use the profile from a user-defined AWS credentials location. |

#### Future work

It occurs to me that this tool could be extended easily enough to act as a Lambda function which would allow messages posted to an S3 bucket to be dispatched to Slack. That's an interesting idea but likely overkill. In any event SFO Museum has no need for this functionality (yet).

## See also

* https://gocloud.dev/howto/runtimevar
* https://github.com/sfomuseum/runtimevar
* https://github.com/aaronland/go-aws-session
