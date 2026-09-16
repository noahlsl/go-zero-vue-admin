package middleware

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"zero/internal/dao/model"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

const maxBodyLogSize = 1024

// OperationLoggerMiddleware 操作记录中间件
type OperationLoggerMiddleware struct {
	db *gorm.DB
}

// NewOperationLoggerMiddleware 创建操作记录中间件
func NewOperationLoggerMiddleware(db *gorm.DB) *OperationLoggerMiddleware {
	return &OperationLoggerMiddleware{db: db}
}

// Handle 记录请求操作日志并写入 sys_operation_records 表
func (m *OperationLoggerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body []byte
		var err error

		if r.Method != http.MethodGet {
			body, err = io.ReadAll(r.Body)
			if err != nil {
				logx.Errorw("读取请求体失败",
					logx.Field("error", err),
					logx.Field("path", r.URL.Path),
				)
			} else {
				r.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			body = marshalQueryParams(r)
		}

		record := buildOperationRecord(r, body)

		writer := &responseBodyWriter{
			ResponseWriter: w,
			body:           &bytes.Buffer{},
		}
		now := time.Now()

		next(writer, r)

		completeRecord(record, writer, now)
		m.saveRecord(r.Context(), record)
	}
}

// marshalQueryParams 将 GET 请求的查询参数序列化为 JSON
func marshalQueryParams(r *http.Request) []byte {
	query := r.URL.RawQuery
	query, _ = url.QueryUnescape(query)
	split := strings.Split(query, "&")
	m := make(map[string]string)
	for _, v := range split {
		kv := strings.SplitN(v, "=", 2)
		if len(kv) == 2 {
			m[kv[0]] = kv[1]
		}
	}
	body, _ := sonic.Marshal(&m)
	return body
}

// buildOperationRecord 构建操作记录（请求阶段）
func buildOperationRecord(r *http.Request, body []byte) *model.SysOperationRecord {
	record := &model.SysOperationRecord{
		Ip:     extractClientIP(r),
		Method: r.Method,
		Path:   r.URL.Path,
		Agent:  r.UserAgent(),
		Body:   truncateBody(r, body),
		UserID: extractUserID(r),
	}
	return record
}

// extractClientIP 提取客户端真实 IP
func extractClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	ip, _, _ := strings.Cut(r.RemoteAddr, ":")
	return ip
}

// truncateBody 截断过大的请求体
func truncateBody(r *http.Request, body []byte) string {
	if strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
		return "[文件]"
	}
	if len(body) > maxBodyLogSize {
		return "[超出记录长度]"
	}
	return string(body)
}

// extractUserID 从请求头或 JWT context 中提取用户 ID
func extractUserID(r *http.Request) int {
	if idStr := r.Header.Get("x-user-id"); idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			return id
		}
	}

	payload := r.Context().Value("payload")
	if payload == nil {
		return 0
	}

	claims, ok := payload.(map[string]interface{})
	if !ok {
		return 0
	}

	// 尝试从 "id" 或 "userId" 字段获取用户 ID
	for _, key := range []string{"id", "userId"} {
		if val, exists := claims[key]; exists {
			return toInt(val)
		}
	}
	return 0
}

// toInt 将 interface{} 转换为 int
func toInt(val interface{}) int {
	switch v := val.(type) {
	case float64:
		return int(v)
	case int:
		return v
	case int64:
		return int(v)
	case uint:
		return int(v)
	default:
		return 0
	}
}

// completeRecord 完成操作记录（响应阶段）
func completeRecord(record *model.SysOperationRecord, writer *responseBodyWriter, start time.Time) {
	record.Latency = time.Since(start)
	record.Status = writer.Status()
	record.Resp = truncateResp(writer.body.String())
}

// truncateResp 截断过大的响应体
func truncateResp(resp string) string {
	if len(resp) > maxBodyLogSize {
		return "[超出记录长度]"
	}
	return resp
}

// saveRecord 异步保存操作记录到数据库
func (m *OperationLoggerMiddleware) saveRecord(ctx context.Context, record *model.SysOperationRecord) {
	if err := m.db.WithContext(ctx).Create(record).Error; err != nil {
		logx.Errorw("创建操作记录失败",
			logx.Field("error", errors.Wrap(err, "写入操作记录")),
			logx.Field("path", record.Path),
			logx.Field("method", record.Method),
		)
	}
}

// responseBodyWriter 包装 http.ResponseWriter 以捕获响应体
type responseBodyWriter struct {
	http.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseBodyWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseBodyWriter) Status() int {
	if w.status != 0 {
		return w.status
	}
	return http.StatusOK
}
