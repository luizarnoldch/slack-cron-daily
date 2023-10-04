package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
)

func handler(ctx context.Context) (string, error) {

	slackToken := os.Getenv("SLACK_TOKEN")
	channel := "#testing"

	limaTimeZone := time.FixedZone("Lima Time", -5*60*60)
	limaTime := time.Now().In(limaTimeZone)

	fmt.Println("Lima Time:", limaTime.Format("2006-01-02 15:04:05.999999999 -0700 MST"))

	morningStart := time.Date(limaTime.Year(), limaTime.Month(), limaTime.Day(), 8, 50, 0, 0, limaTimeZone)
	morningEnd := time.Date(limaTime.Year(), limaTime.Month(), limaTime.Day(), 9, 20, 0, 0, limaTimeZone)

	fmt.Println("Morning Start:", morningStart)
	fmt.Println("Morning End:", morningEnd)

	lunchStart := time.Date(limaTime.Year(), limaTime.Month(), limaTime.Day(), 12, 50, 0, 0, limaTimeZone)
	lunchBreak := time.Date(limaTime.Year(), limaTime.Month(), limaTime.Day(), 13, 20, 0, 0, limaTimeZone)

	fmt.Println("Lunch Start:", lunchStart)
	fmt.Println("Lunch Break:", lunchBreak)

	lunchBackStart := time.Date(limaTime.Year(), limaTime.Month(), limaTime.Day(), 13, 50, 0, 0, limaTimeZone)
	lunchBackEnd := time.Date(limaTime.Year(), limaTime.Month(), limaTime.Day(), 14, 10, 0, 0, limaTimeZone)

	fmt.Println("Lunch Back Start:", lunchBackStart)
	fmt.Println("Lunch Back End:", lunchBackEnd)

	eveningStart := time.Date(limaTime.Year(), limaTime.Month(), limaTime.Day(), 18, 0, 0, 0, limaTimeZone)
	eveningEnd := time.Date(limaTime.Year(), limaTime.Month(), limaTime.Day(), 18, 15, 0, 0, limaTimeZone)

	fmt.Println("Evening Start:", eveningStart)
	fmt.Println("Evening End:", eveningEnd)

	var message string
	if limaTime.After(morningStart) && limaTime.Before(morningEnd) {
		message = "/wfh hi"
	} else if limaTime.After(lunchStart) && limaTime.Before(lunchBreak) {
		message = "/wfh break"
	} else if limaTime.After(lunchBackStart) && limaTime.Before(lunchBackEnd) {
		message = "/wfh back"
	} else if limaTime.After(eveningStart) && limaTime.Before(eveningEnd) {
		message = "/wfh bye"
	}

	if message != "" {
		api := slack.New(slackToken)
		_, _, err := api.PostMessage(channel, slack.MsgOptionText(message, false))
		if err != nil {
			log.Printf("Error sending message to Slack: %v", err)
			return "ERROR", nil
		}
		fmt.Printf("Sent message: %s\n", message)
	} else {
		fmt.Println("No scheduled message at this time.")
	}

	fmt.Println("SUCCESS")
	return "SUCCESS", nil
}

func main() {
	lambda.Start(handler)
}
