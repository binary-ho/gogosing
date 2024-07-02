package oauth

import (
	"gogosing/server/goauth/util"
)

type oAuthProperty struct {
	OAuth struct {
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
		RedirectURL  string `yaml:"redirect_url"`
		Endpoint     string `yaml:"endpoint"`
	}
}

const (
	oAuthPropertyPath = "server\\goauth\\secret\\property.yml"
)

func init() {
	oAuthProperty := &oAuthProperty{}
	err := util.FetchYAML(oAuthPropertyPath, &oAuthProperty)
	InitOAuthConfig(oAuthProperty)
	if err != nil {
		panic(err)
	}
}
