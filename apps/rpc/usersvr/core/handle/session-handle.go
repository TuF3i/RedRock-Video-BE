package handle

import (
	"LiveDanmu/apps/public/jwt"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/usersvr/core"
	"LiveDanmu/apps/rpc/usersvr/core/dto"
	"LiveDanmu/apps/rpc/usersvr/core/pkg"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
	"context"

	"go.uber.org/zap"
)

func convertRvUserInfo2RvUser(raw *usersvr.RvUserInfo) *dao.RvUser {
	bio := ""
	if raw.Bio != nil {
		bio = *raw.Bio
	}
	return &dao.RvUser{
		Uid:       raw.Uid,
		Login:     raw.UserName,
		AvatarURL: raw.AvatarUrl,
		Bio:       bio,
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
	uid := uInfo.GetUid()

	core.Logger.INFO("UserLogin start", zap.Int64("uid", uid), zap.String("user_name", uInfo.GetUserName()))

	if !pkg.ValidateGitHubUserID(uid) {
		core.Logger.WARN("UserLogin invalid uid", zap.Int64("uid", uid))
		return dto.InvalidUID, nil
	}
	if !pkg.ValidateGitHubUserLogin(uInfo.GetUserName()) {
		core.Logger.WARN("UserLogin invalid user name", zap.Int64("uid", uid), zap.String("user_name", uInfo.GetUserName()))
		return dto.InvalidUserName, nil
	}
	if !pkg.ValidateGitHubUserAvatarURL(uInfo.GetAvatarUrl()) {
		core.Logger.WARN("UserLogin invalid avatar url", zap.Int64("uid", uid), zap.String("avatar_url", uInfo.GetAvatarUrl()))
		return dto.InvalidAvatarURL, nil
	}
	if !pkg.ValidateGitHubUserBio(uInfo.GetBio()) {
		core.Logger.WARN("UserLogin invalid bio", zap.Int64("uid", uid))
		return dto.InvalidBio, nil
	}

	data := convertRvUserInfo2RvUser(uInfo)
	err := core.Dao.AddUser(data)
	if err != nil {
		core.Logger.WARN("UserLogin add user failed", zap.Int64("uid", uid), zap.Error(err))
		return dto.ServerInternalError(err), nil
	}

	u, err := core.Dao.GetUserInfo(uid)
	if err != nil {
		core.Logger.WARN("UserLogin get user info failed", zap.Int64("uid", uid), zap.Error(err))
		return dto.ServerInternalError(err), nil
	}

	role := u.Role
	accessToken, err := jwt.GenerateAccessToken(uid, role)
	if err != nil {
		core.Logger.WARN("UserLogin generate access token failed", zap.Int64("uid", uid), zap.Error(err))
		return dto.ServerInternalError(err), nil
	}

	refreshToken, err := jwt.GenerateRefreshToken(uid, role)
	if err != nil {
		core.Logger.WARN("UserLogin generate refresh token failed", zap.Int64("uid", uid), zap.Error(err))
		return dto.ServerInternalError(err), nil
	}

	err = core.Dao.SetNewAccessToken(ctx, uid, accessToken)
	if err != nil {
		core.Logger.WARN("UserLogin set access token failed", zap.Int64("uid", uid), zap.Error(err))
		return dto.ServerInternalError(err), nil
	}
	err = core.Dao.SetNewRefreshToken(ctx, uid, refreshToken)
	if err != nil {
		core.Logger.WARN("UserLogin set refresh token failed", zap.Int64("uid", uid), zap.Error(err))
		return dto.ServerInternalError(err), nil
	}

	loginData := genLoginData(accessToken, refreshToken)

	core.Logger.INFO("UserLogin success", zap.Int64("uid", uid), zap.String("role", role))
	return dto.OperationSuccess, loginData
}

func UserLogout(ctx context.Context, req *usersvr.LogoutReq) dto.Response {
	uid := req.GetUid()

	core.Logger.INFO("UserLogout start", zap.Int64("uid", uid))

	err := core.Dao.Logout(ctx, uid)
	if err != nil {
		core.Logger.WARN("UserLogout failed", zap.Int64("uid", uid), zap.Error(err))
		return dto.ServerInternalError(err)
	}

	core.Logger.INFO("UserLogout success", zap.Int64("uid", uid))
	return dto.OperationSuccess
}
