package main

import (
	usersvr "LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
	"context"
)

// UserSvrImpl implements the last service interface defined in the IDL.
type UserSvrImpl struct{}

// UserLogin implements the UserSvrImpl interface.
func (s *UserSvrImpl) UserLogin(ctx context.Context, req *usersvr.LoginReq) (resp *usersvr.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// RefreshToken implements the UserSvrImpl interface.
func (s *UserSvrImpl) RefreshToken(ctx context.Context, req *usersvr.RefreshReq) (resp *usersvr.RefreshResp, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserSvrImpl interface.
func (s *UserSvrImpl) GetUserInfo(ctx context.Context, req *usersvr.GetUserInfoReq) (resp *usersvr.GetUserInfoResp, err error) {
	// TODO: Your code here...
	return
}

// SetAdminRole implements the UserSvrImpl interface.
func (s *UserSvrImpl) SetAdminRole(ctx context.Context, req *usersvr.SetAdminRoleReq) (resp *usersvr.SetAdminRoleResp, err error) {
	// TODO: Your code here...
	return
}

// GetAdminer implements the UserSvrImpl interface.
func (s *UserSvrImpl) GetAdminer(ctx context.Context, req *usersvr.GetAdminerReq) (resp *usersvr.GetAdminerResp, err error) {
	// TODO: Your code here...
	return
}

// GetUsers implements the UserSvrImpl interface.
func (s *UserSvrImpl) GetUsers(ctx context.Context, req *usersvr.GetUsersReq) (resp *usersvr.GetUsersResp, err error) {
	// TODO: Your code here...
	return
}

// Logout implements the UserSvrImpl interface.
func (s *UserSvrImpl) Logout(ctx context.Context, req *usersvr.LogoutReq) (resp *usersvr.LogoutResp, err error) {
	// TODO: Your code here...
	return
}
