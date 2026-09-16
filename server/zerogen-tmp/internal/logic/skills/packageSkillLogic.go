// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package skills

import (
	"context"

	"zero/zerogen-tmp/internal/svc"
	"zero/zerogen-tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PackageSkillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPackageSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PackageSkillLogic {
	return &PackageSkillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PackageSkillLogic) PackageSkill(req *types.PackageSkillReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
