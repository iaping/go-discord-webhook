package main

import (
	"log"

	webhook "github.com/iaping/go-discord-webhook"
)

func main() {
	hook := webhook.New("https://discord.com/api/webhooks/xxxxxxxxxx")

	// say
	if err := hook.Say(":grin: Hello,World!"); err != nil {
		log.Println(err)
	}

	// send
	if err := hook.Send(webhook.Message{
		Content: "Hello,World!",
		Embeds: []webhook.Embed{
			{
				Title:       "Title",
				Description: "Description",
				Url:         "https://www.google.com",
				Image: &webhook.Media{
					Url: "https://ik.imagekit.io/cymbal/6387ec220e795816d4b70984?tr=w-500",
				},
				Footer: &webhook.Footer{
					Text: "Footer",
				},
				Fields: []webhook.Field{
					{
						Name:   ":grin:field1",
						Value:  "value1",
						Inline: true,
					},
					{
						Name:   "field2",
						Value:  "value2",
						Inline: true,
					},
				},
			},
		},
	}); err != nil {
		log.Println(err)
	}
}
