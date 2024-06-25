package oauth

import (
	"encoding/base64"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"math/rand"
)

var Config *oauth2.Config

func InitOAuthConfig(property *oAuthProperty) {
	Config = &oauth2.Config{
		ClientID:     property.OAuth.ClientID,
		ClientSecret: property.OAuth.ClientSecret,
		RedirectURL:  property.OAuth.RedirectURL,
		Endpoint:     github.Endpoint,
	}
}

func GetRandomLoginURL() string {
	token := RandToken()
	return Config.AuthCodeURL(token)
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
