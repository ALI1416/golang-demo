package logic

import (
	"context"

	"test/demo/internal/svc"
	"test/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (resp []types.ListItem, err error) {
	resp = *new([]types.ListItem)
	resp = append(resp, types.ListItem{Value: "aaa"}, types.ListItem{Value: "bbb"})
	return
}
