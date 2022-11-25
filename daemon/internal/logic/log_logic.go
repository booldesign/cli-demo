package logic

import (
	"context"

	"github.com/booldesign/cli-demo/daemon/internal/svc"
	model "github.com/booldesign/cli-demo/model/usercenter"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 22:30
 * @Desc:
 */

type InsertLogLogic struct {
	svcCtx *svc.ServiceContext
}

func NewInsertLogLogic(svcCtx *svc.ServiceContext) *InsertLogLogic {
	return &InsertLogLogic{
		svcCtx: svcCtx,
	}
}

func (l *InsertLogLogic) InsertLog(ctx context.Context, data *model.LogModel) error {
	return l.svcCtx.Db.Save(&data).Error
}
