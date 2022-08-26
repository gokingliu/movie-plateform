package services

import (
	"context"
	pb "git.woa.com/crotaliu/pb-hub"
)

type InfoImpl struct{}

// GetInfo 获取视频详情
func (s *InfoImpl) GetInfo(ctx context.Context, req *pb.GetInfoReq, rsp *pb.GetInfoRsp) error {

	return nil
}
