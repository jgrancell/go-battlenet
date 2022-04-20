package blizzapi

import (
	"time"

	"github.com/jgrancell/go-battlenet/token"
)

type Client struct {
	Config    Config
	SubClient string
	Token     *token.Token
}

func (c *Client) Init() *Client {
	c.Config.Init()

	c.Token = &token.Token{
		ClientID:     c.Config.ClientID,
		ClientSecret: c.Config.ClientSecret,
		Expiration:   time.Now(),
		Region:       c.Config.Region,
	}
	c.Token.Renew()

	return c
}

func (c *Client) Wow() *Client {
	c.Init()
	c.SubClient = "wow"
	return c
}
