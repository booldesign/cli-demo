package server

import (
	"context"

	"github.com/booldesign/cli-demo/daemon/internal/logic"
	"github.com/booldesign/cli-demo/daemon/internal/svc"
	"github.com/booldesign/cli-demo/model/usercenter"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 22:36
 * @Desc:
 */

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) InsertLog(ctx context.Context, in *usercenter.LogModel) error {
	l := logic.NewInsertLogLogic(s.svcCtx)
	return l.InsertLog(ctx, in)
}
