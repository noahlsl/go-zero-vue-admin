package sysDictionaryDetail

import (
	"context"
	"strconv"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysDictionaryDetailLogic {
	return &UpdateSysDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysDictionaryDetailLogic) UpdateSysDictionaryDetail(req *types.UpdateSysDictionaryDetailReq) (resp *types.Response, err error) {
	// 查询原记录
	var detail model.SysDictionaryDetail
	if err = l.svcCtx.DB.WithContext(l.ctx).First(&detail, req.ID).Error; err != nil {
		return nil, errors.Wrap(err, "查询字典详情失败")
	}

	// 更新 ParentID 时重新计算层级和路径
	if req.ParentID != nil {
		var parent model.SysDictionaryDetail
		if err = l.svcCtx.DB.WithContext(l.ctx).First(&parent, *req.ParentID).Error; err != nil {
			return nil, errors.Wrap(err, "查询父级字典详情失败")
		}

		// 检查循环引用
		if l.checkCircularReference(req.ID, *req.ParentID) {
			return nil, errors.New("不能将字典详情设置为自己或其子项的父级")
		}

		detail.ParentID = req.ParentID
		detail.Level = parent.Level + 1
		if parent.Path == "" {
			detail.Path = strconv.Itoa(int(parent.ID))
		} else {
			detail.Path = parent.Path + "," + strconv.Itoa(int(parent.ID))
		}
	} else {
		detail.ParentID = nil
		detail.Level = 0
		detail.Path = ""
	}

	detail.Label = req.Label
	detail.Value = req.Value
	detail.Extend = req.Extend
	detail.Status = req.Status
	detail.Sort = req.Sort

	if err = l.svcCtx.DB.WithContext(l.ctx).Save(&detail).Error; err != nil {
		return nil, errors.Wrap(err, "更新字典详情失败")
	}

	// 更新子项的层级和路径
	if err = l.updateChildrenLevelAndPath(detail.ID); err != nil {
		return nil, err
	}

	return &types.Response{Code: 0, Msg: "修改成功"}, nil
}

// checkCircularReference 检查循环引用
func (l *UpdateSysDictionaryDetailLogic) checkCircularReference(id, parentID uint) bool {
	if id == parentID {
		return true
	}

	var parent model.SysDictionaryDetail
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&parent, parentID).Error; err != nil {
		return false
	}
	if parent.ParentID == nil {
		return false
	}
	return l.checkCircularReference(id, *parent.ParentID)
}

// updateChildrenLevelAndPath 更新子项的层级和路径
func (l *UpdateSysDictionaryDetailLogic) updateChildrenLevelAndPath(parentID uint) error {
	var children []model.SysDictionaryDetail
	if err := l.svcCtx.DB.WithContext(l.ctx).
		Where("parent_id = ?", parentID).Find(&children).Error; err != nil {
		return errors.Wrap(err, "查询子项失败")
	}

	var parent model.SysDictionaryDetail
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&parent, parentID).Error; err != nil {
		return errors.Wrap(err, "查询父级失败")
	}

	for _, child := range children {
		child.Level = parent.Level + 1
		if parent.Path == "" {
			child.Path = strconv.Itoa(int(parent.ID))
		} else {
			child.Path = parent.Path + "," + strconv.Itoa(int(parent.ID))
		}

		if err := l.svcCtx.DB.WithContext(l.ctx).Save(&child).Error; err != nil {
			return errors.Wrap(err, "更新子项失败")
		}

		// 递归更新子项的子项
		if err := l.updateChildrenLevelAndPath(child.ID); err != nil {
			return err
		}
	}

	return nil
}
