package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// This class holds all the user data we get from the Discord API
type DiscordAccount struct {
	Token					string		`json:"token"`
	ID						string		`json:"id"`
	Username				string		`json:"username"`
	Avatar					string		`json:"avatar"`
	Discriminator			string		`json:"discriminator"`
	PublicFlags				int			`json:"public_flags"`
	PremiumType				int			`json:"premium_type"`
	Flags					int			`json:"flags"`
	Banner					bool		`json:"banner"`
	AccentColor				int			`json:"accent_color"`
	GlobalName				string		`json:"global_name"`
	AvatarDecorationData	bool		`json:"avatar_decoration_data"`
	BannerColor				string		`json:"banner_color"`
	MFAEnabled				bool		`json:"mfa_enabled"`
	Locale					string		`json:"locale"`
	Email					string		`json:"email"`
	Verified				bool		`json:"verified"`
	Phone					string		`json:"phone"`
	NSFWAllowed				bool		`json:"nsfw_allowed"`
	PremiumUsageFlags		int			`json:"premium_usage_flags"`
	LinkedUsers				[]string	`json:"linked_users"`
	PurchasedFlags			int			`json:"purchased_flags"`
	Bio						string		`json:"bio"`
	AuthenticatorTypes		[]int		`json:"authenticator_types"`
}

func getAccountData(token string) (account DiscordAccount, err error) {
    // Create a new GET request to the Discord API
    request, err := http.NewRequest("GET", "https://discordapp.com/api/v9/users/@me", nil)
    	if err != nil {return DiscordAccount{}, err}

    // Set proper request headers
	request.Header.Set("Authorization", token)
	request.Header.Set("Content-Type", "application/json")
	// Setting user-agent to curl makes the reponse much prettier - don't have to clean/parse it at all
    request.Header.Set("User-Agent", "curl/7.37.0")

    // Send the request
    client := http.Client{}
    response, err := client.Do(request)
    	if err != nil {return DiscordAccount{}, err}
    defer response.Body.Close() // Make sure we release resources associated with the response

	// Do some error checking on the reponse code. Discord API responds with 401 when
	// token is invalid but we might as well check for anything that isn't 200
	if response.StatusCode != http.StatusOK {
		return DiscordAccount{}, fmt.Errorf("response code %v is invalid", response.StatusCode)
	}

    // Decode the JSON data we got from the API and read it into an instance of DiscordAccount struct.
	// This is pretty simple because the struct is built based off what the API returns, so we
	// can just feed the data right back into the DiscordAccount instance without much effort
    err = json.NewDecoder(response.Body).Decode(&account)
    	if err != nil {return DiscordAccount{}, err}

    // Return account data if no errors encountered
	return account, nil
}