/*
Create time at 2023年3月11日0011下午 14:01:02
Create User at Administrator
*/

package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"os"
)

func main() {
	os.Setenv("socket_bot_token", "xoxb-4947141615057-4931591814133-0oAYyLKXYKfmXUFezJisDVIp")
	os.Setenv("channel_id", "C04TDH11Y9K")
	api := slack.New(os.Getenv("socket_bot_token"))
	channel := []string{os.Getenv("channel_id")}
	fileadr := []string{"test.png"}

	for i := 0; i < len(fileadr); i++ {
		params := slack.FileUploadParameters{
			Channels: channel,
			File:     fileadr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("file.Name: %s,file.Url: %s\n", file.Name, file.URL)
	}
}
