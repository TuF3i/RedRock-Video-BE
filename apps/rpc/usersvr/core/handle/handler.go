package handle

import (
	"LiveDanmu/apps/rpc/usersvr/core/dto"
	usersvr "LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
	"context"
	"errors"
)

// UserSvrImpl implements the last service interface defined in the IDL.
type UserSvrImpl struct{}

// UserLogin implements the UserSvrImpl interface.
func (s *UserSvrImpl) UserLogin(ctx context.Context, req *usersvr.LoginReq) (resp *usersvr.LoginResp, err error) {
	// 调用方法
	rawResp, data := UserLogin(ctx, req)
	resp = dto.GenKitexResp[*usersvr.LoginResp](rawResp, data)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// RefreshToken implements the UserSvrImpl interface.
func (s *UserSvrImpl) RefreshToken(ctx context.Context, req *usersvr.RefreshReq) (resp *usersvr.RefreshResp, err error) {
	// 调用方法
	rawResp, data := GetRefreshToken(ctx, req)
	resp = dto.GenKitexResp[*usersvr.RefreshResp](rawResp, data)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// GetUserInfo implements the UserSvrImpl interface.
func (s *UserSvrImpl) GetUserInfo(ctx context.Context, req *usersvr.GetUserInfoReq) (resp *usersvr.GetUserInfoResp, err error) {
	// 调用方法
	rawResp, data := GetUserInfo(ctx, req)
	resp = dto.GenKitexResp[*usersvr.GetUserInfoResp](rawResp, data)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// SetAdminRole implements the UserSvrImpl interface.
func (s *UserSvrImpl) SetAdminRole(ctx context.Context, req *usersvr.SetAdminRoleReq) (resp *usersvr.SetAdminRoleResp, err error) {
	// 调用方法
	rawResp := SetAdminRole(ctx, req)
	resp = dto.GenKitexResp[*usersvr.SetAdminRoleResp](rawResp, nil)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// GetAdminer implements the UserSvrImpl interface.
func (s *UserSvrImpl) GetAdminer(ctx context.Context, req *usersvr.GetAdminerReq) (resp *usersvr.GetAdminerResp, err error) {
	// 调用方法
	rawResp, data := GetAdminList(ctx, req)
	resp = dto.GenKitexResp[*usersvr.GetAdminerResp](rawResp, data)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// GetUsers implements the UserSvrImpl interface.
func (s *UserSvrImpl) GetUsers(ctx context.Context, req *usersvr.GetUsersReq) (resp *usersvr.GetUsersResp, err error) {
	// 调用方法
	rawResp, data := GetUserList(ctx, req)
	resp = dto.GenKitexResp[*usersvr.GetUsersResp](rawResp, data)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// Logout implements the UserSvrImpl interface.
func (s *UserSvrImpl) Logout(ctx context.Context, req *usersvr.LogoutReq) (resp *usersvr.LogoutResp, err error) {
	// 调用方法
	rawResp := UserLogout(ctx, req)
	resp = dto.GenKitexResp[*usersvr.LogoutResp](rawResp, nil)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}
