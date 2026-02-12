package OAuth2

import (
	"LiveDanmu/apps/public/config/config_template"

	"golang.org/x/oauth2"
)

type OAuthCore struct {
	conf      *config_template.UserGatewayConfig
	oauthConf *oauth2.Config
}
