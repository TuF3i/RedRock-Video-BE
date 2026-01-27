package pub

import (
	pub "LiveDanmu/apps/rpc/pub/kitex_gen/pub"
	"context"
)

// PubImpl implements the last service interface defined in the IDL.
type PubImpl struct{}

// Pub implements the PubImpl interface.
func (s *PubImpl) Pub(ctx context.Context, req *pub.PubReq) (resp *pub.PubResp, err error) {
	// TODO: Your code here...
	return
}
