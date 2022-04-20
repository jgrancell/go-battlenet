package profile

import (
	"github.com/jgrancell/go-battlenet/token"
)

type CharacterClient struct {
	BaseUrl   string
	Locale    string
	Namespace string
	Token     *token.Token
}
