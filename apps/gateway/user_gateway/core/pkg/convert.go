package pkg

import (
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
	"LiveDanmu/apps/shared/models"
)

func ConvertGitHubUser2RvUserInfo(raw *models.GitHubUser) *usersvr.RvUserInfo {
	bio := raw.Bio
	role := ""
	return &usersvr.RvUserInfo{
		Uid:       raw.ID,
		UserName:  raw.Login,
		AvatarUrl: raw.AvatarURL,
		Bio:       &bio,
		Role:      &role,
	}
}
