package pkg

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
)

func ConvertGitHubUser2RvUserInfo(raw *dao.GitHubUser) *usersvr.RvUserInfo {
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
