package blizzapi

import (
	"fmt"

	"github.com/jgrancell/go-battlenet/wow/profile"
)

func (c *Client) Character(realm string, name string) *profile.CharacterClient {
	url := fmt.Sprintf(
		"https://%s/%s/%s/%s",
		c.Config.BaseUrl,
		"profile/wow/character",
		realm,
		name,
	)
	char := &profile.CharacterClient{
		BaseUrl:   url,
		Locale:    c.Config.Locale,
		Namespace: "profile-" + c.Config.Region,
		Token:     c.Token,
	}

	return char
}
