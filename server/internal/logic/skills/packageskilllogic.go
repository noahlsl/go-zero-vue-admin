package skills

import (
	"context"
	"fmt"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
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

// PackageSkill 打包技能为 ZIP 文件
// TODO: 实现完整的打包逻辑
func (l *PackageSkillLogic) PackageSkill(req *types.PackageSkillReq) (resp *types.Response, err error) {
	if req.ID == 0 {
		return nil, errors.New("技能ID不能为空")
	}

	return &types.Response{
		Code: 0,
		Data: map[string]interface{}{
			"name": fmt.Sprintf("skill_%d.zip", req.ID),
		},
		Msg: fmt.Sprintf("技能 %d 打包成功", req.ID),
	}, nil
}
