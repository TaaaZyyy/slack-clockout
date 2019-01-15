package main

import (
	"time"
	"context"
	"os"

	slack "./slack"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	DeviceEvent struct {
		ButtonClicked struct {
			ClickType string `json:"clickType"`
		} `json:"buttonClicked"`
	} `json:"deviceEvent"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	t := time.Now()
	nowUTC := t.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	const layout = "2006-01-02 15:04:05"
	now := nowJST.Format(layout)

	slackURL  := os.Getenv("SlackURL")
	slackName := os.Getenv("Name")
	lastName  := os.Getenv("LastName")
	privatChannel := os.Getenv("PrivateChannel")
	publicChannel := os.Getenv("PublicChannel")


	clicktype := event.DeviceEvent.ButtonClicked.ClickType

	if clicktype == "SINGLE" {
		sl := slack.NewSlack(slackURL, "退勤時間: " + now, slackName, privatChannel)
		sl.Send()
	} else if clicktype == "DOUBLE" {
		sl := slack.NewSlack(slackURL, "最終退勤者: " + lastName, slackName, publicChannel)
		sl.Send()
	} else if clicktype == "LONG" {
		// sl := slack.NewSlack(slackURL, "ロングクリック", "", "")
		// sl.Send()
	}
	return "", nil
}

func main() {
	lambda.Start(HandleRequest)
}