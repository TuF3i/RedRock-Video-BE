package OAuth2

import (
	"LiveDanmu/apps/public/config/config_template"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type OAuthCore struct {
	conf      *config_template.UserGatewayConfig
	oauthConf *oauth2.Config
}

func GetOAuth2(conf *config_template.UserGatewayConfig) *OAuthCore {
	o := &OAuthCore{
		conf: conf,
		oauthConf: &oauth2.Config{
			ClientID:     conf.OAuth.ClientID,
			ClientSecret: conf.OAuth.ClientSecret,
			Endpoint:     github.Endpoint,
			RedirectURL:  conf.OAuth.RedirectURL,
			Scopes:       []string{},
		},
	}

	return o
}
