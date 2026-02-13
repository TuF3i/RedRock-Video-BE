package handle

import (
	"LiveDanmu/apps/public/jwt"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/usersvr/core"
	"LiveDanmu/apps/rpc/usersvr/core/dto"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
	"context"
)

func convertRvUserInfo2RvUser(raw *usersvr.RvUserInfo) *dao.RvUser {
	return &dao.RvUser{
		Uid:       raw.Uid,
		Login:     raw.UserName,
		AvatarURL: raw.AvatarUrl,
		Bio:       raw.Bio,
	}
}

func genLoginData(accessToken string, refreshToken string) *usersvr.LoginData {
	return &usersvr.LoginData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func UserLogin(ctx context.Context, req *usersvr.LoginReq) (dto.Response, *usersvr.LoginData) {
	// 获取UserInfo
	uInfo := req.GetUserInfo()
	// 转换结构体
	data := convertRvUserInfo2RvUser(uInfo)
	// 将用户写入数据库
	err := core.Dao.AddUser(data)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 获取用户信息
	u, err := core.Dao.GetUserInfo(uInfo.GetUid())
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 生成AccessToken
	role := u.Role
	uid := u.Uid
	accessToken, err := jwt.GenerateAccessToken(uid, role)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 生成refreshToken
	refreshToken, err := jwt.GenerateRefreshToken(uid, role)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 写入redis
	err = core.Dao.SetNewAccessToken(ctx, uid, accessToken)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}
	err = core.Dao.SetNewRefreshToken(ctx, uid, refreshToken)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 组装data
	loginData := genLoginData(accessToken, refreshToken)

	return dto.OperationSuccess, loginData
}
