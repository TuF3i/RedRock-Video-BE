package handle

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
