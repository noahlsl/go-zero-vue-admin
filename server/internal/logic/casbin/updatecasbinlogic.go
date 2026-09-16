package casbin

import (
	"context"
	"strconv"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCasbinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCasbinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCasbinLogic {
	return &UpdateCasbinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCasbinLogic) UpdateCasbin(req *types.UpdateCasbinReq) (resp *types.Response, err error) {
	authorityId := strconv.Itoa(int(req.AuthorityId))

	// 1. 删除该角色原有的所有策略
	if err = l.clearCasbinByAuthorityId(authorityId); err != nil {
		return nil, errors.Wrap(err, "清除角色旧策略失败")
	}

	// 2. 去重后批量插入新策略
	rules := l.deduplicateCasbinInfos(authorityId, req.CasbinInfos)
	if len(rules) == 0 {
		return &types.Response{Code: 0, Msg: "success"}, nil
	}

	if err = l.batchInsertCasbinRules(rules); err != nil {
		return nil, errors.Wrap(err, "批量插入Casbin策略失败")
	}

	logx.WithContext(l.ctx).Infow("更新Casbin策略成功",
		logx.Field("authority_id", authorityId),
		logx.Field("policy_count", len(rules)),
		logx.Field("module", "casbin"),
		logx.Field("action", "update_casbin"),
	)

	return &types.Response{Code: 0, Msg: "success"}, nil
}

// clearCasbinByAuthorityId 根据角色ID清除所有策略
func (l *UpdateCasbinLogic) clearCasbinByAuthorityId(authorityId string) error {
	return l.svcCtx.DB.WithContext(l.ctx).
		Where("ptype = ? AND v0 = ?", "p", authorityId).
		Delete(&model.SysCasbinRule{}).
		Error
}

// deduplicateCasbinInfos 对策略列表进行去重处理
func (l *UpdateCasbinLogic) deduplicateCasbinInfos(authorityId string, casbinInfos []types.CasbinInfo) []model.SysCasbinRule {
	deduplicateMap := make(map[string]bool)
	var rules []model.SysCasbinRule

	for _, info := range casbinInfos {
		key := authorityId + info.Path + info.Method
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			rules = append(rules, model.SysCasbinRule{
				Ptype: "p",
				V0:    authorityId,
				V1:    info.Path,
				V2:    info.Method,
			})
		}
	}

	return rules
}

// batchInsertCasbinRules 批量插入策略规则
func (l *UpdateCasbinLogic) batchInsertCasbinRules(rules []model.SysCasbinRule) error {
	return l.svcCtx.DB.WithContext(l.ctx).Create(&rules).Error
}
