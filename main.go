package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/acarl005/stripansi"
)

var wg sync.WaitGroup

func main() {
	var msg;
	msg := data{
	    MsgType: "text",
	    Content: struct {
		Text string `json:"text"`
	    }{Text: stripansi.Strip(line)},
	}
	var oneLine, verboseMode bool
	var webhookURL, lines string
	flag.StringVar(&webhookURL, "u", "", "Slack Webhook URL")
	flag.BoolVar(&oneLine, "1", false, "Send message line-by-line")
	flag.BoolVar(&verboseMode, "v", false, "Verbose mode")
	flag.Parse()

	webhookEnv := os.Getenv("SLACK_WEBHOOK_URL")
	if webhookEnv != "" {
		webhookURL = webhookEnv
	} else {
		if webhookURL == "" {
			if verboseMode {
				fmt.Println("Slack Webhook URL not set!")
			}
		}
	}

	if !isStdin() {
		os.Exit(1)
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()

		fmt.Println(line)
		if oneLine {
			if webhookURL != "" {
				wg.Add(1)
				go slackCat(webhookURL, line)
			}
		} else {
			lines += line
			lines += "\n"
		}
	}

	if !oneLine {
		wg.Add(1)
		go slackCat(webhookURL, lines)
	}
	wg.Wait()
}

func isStdin() bool {
	f, e := os.Stdin.Stat()
	if e != nil {
		return false
	}

	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}

type data struct {
    MsgType string `json:"msg_type"`
    Content struct {
        Text string `json:"text"`
    } `json:"content"`
}


func slackCat(url string, line string) {
	data, _ := json.Marshal(msg)
	http.Post(url, "application/json", strings.NewReader(string(data)))
	wg.Done()
}
