package file

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"zero/internal/dao/model"
	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	maxUploadSize = 100 << 20 // 100 MB
)

type httpRequestKeyType struct{}

var httpRequestKey = httpRequestKeyType{}

// WithHTTPRequest 将 HTTP 请求存入上下文，供 logic 层读取
func WithHTTPRequest(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, httpRequestKey, r)
}

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile() (resp *types.ExaFileRes, err error) {
	// 1. 从上下文获取 HTTP 请求
	r, ok := l.ctx.Value(httpRequestKey).(*http.Request)
	if !ok {
		return nil, errors.New("获取HTTP请求失败")
	}

	// 2. 解析 multipart 表单
	if err = r.ParseMultipartForm(maxUploadSize); err != nil {
		return nil, errors.Wrap(err, "解析上传表单失败")
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, errors.Wrap(err, "读取上传文件失败")
	}
	defer file.Close()

	// 3. 解析上传参数
	noSave := r.FormValue("noSave")
	classId := 0
	if v := r.FormValue("classId"); v != "" {
		_, _ = fmt.Sscanf(v, "%d", &classId)
	}

	// 4. 构建文件记录
	record := l.buildFileRecord(header, classId)

	// 5. 仅在 noSave == "0" 时持久化到数据库
	if noSave == "0" {
		if err = l.saveFileRecord(record); err != nil {
			return nil, err
		}
	}

	return &types.ExaFileRes{
		File: l.toResponse(record),
	}, nil
}

// buildFileRecord 根据上传文件头构建数据库记录
func (l *UploadFileLogic) buildFileRecord(header *multipart.FileHeader, classId int) model.ExaFileUploadAndDownload {
	name := header.Filename
	ext := ""
	if idx := strings.LastIndex(name, "."); idx > 0 {
		ext = name[idx+1:]
	}
	key := fmt.Sprintf("%d_%s", header.Size, name)

	return model.ExaFileUploadAndDownload{
		Name:    name,
		Url:     "",
		ClassId: classId,
		Tag:     ext,
		Key:     key,
	}
}

// saveFileRecord 将文件记录保存到数据库（幂等：相同 key 跳过）
func (l *UploadFileLogic) saveFileRecord(record model.ExaFileUploadAndDownload) error {
	var existing model.ExaFileUploadAndDownload
	if err := l.svcCtx.DB.WithContext(l.ctx).Where("key = ?", record.Key).First(&existing).Error; err == nil {
		return nil // 已存在，跳过
	}

	if err := l.svcCtx.DB.WithContext(l.ctx).Create(&record).Error; err != nil {
		l.Logger.Errorw("保存文件记录失败",
			logx.Field("error", err),
			logx.Field("fileName", record.Name),
			logx.Field("module", "file"),
			logx.Field("action", "upload"),
		)
		return errors.Wrap(err, "保存文件记录失败")
	}
	return nil
}

// toResponse 将 model 转换为响应类型
func (l *UploadFileLogic) toResponse(m model.ExaFileUploadAndDownload) types.ExaFileUploadAndDownload {
	return types.ExaFileUploadAndDownload{
		ID:        m.ID,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
		Name:      m.Name,
		Url:       m.Url,
	}
}
