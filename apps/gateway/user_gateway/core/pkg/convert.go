package pkg

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
)

func ConvertGitHubUser2RvUserInfo(raw *dao.GitHubUser) *usersvr.RvUserInfo {
	return &usersvr.RvUserInfo{
		Uid:       raw.ID,
		UserName:  raw.Login,
		AvatarUrl: raw.AvatarURL,
		Bio:       raw.Bio,
		Role:      "",
	}
}
