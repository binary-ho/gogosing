package oauth

import (
	"fmt"
	"gogosing/goauth/util"
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
	oAuthPropertyPath = "oauth\\secret\\property.yml.yml.yml"
)

func init() {
	fmt.Println("init oAuthProperty")
	oAuthProperty := &oAuthProperty{}
	err := util.FetchYAML(oAuthPropertyPath, &oAuthProperty)
	InitOAuthConfig(oAuthProperty)
	if err != nil {
		panic(err)
	}
}
