#改自slackcat-> Feishucat

A simple way of sending messages from the CLI output to your Feishu with webhook.

## Installation

- If you have go1.13+ compiler installed: `go build -o feishucat  main.go`

## Configuration

**1** _(optional)_**:** Set `FEISHU_WEBHOOK_URL` environment variable.
```bash
export FEISHU_WEBHOOK_URL="https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxxx"
```

## Usage

It's very simple!

```bash
▶ echo -e "Hello,\nworld!" | ./feishucat
```

### Flags

```
Usage of slackcat:
  -1    Send message line-by-line
  -u string
        Slack Webhook URL
  -v    Verbose mode
```

### Workaround

The goal is to get automated alerts for interesting stuff!

```bash
▶ echo 123| feishucat -1
```

