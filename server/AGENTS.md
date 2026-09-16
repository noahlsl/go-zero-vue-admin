# AGENTS.md — go-zero 后端项目

## 目的

本文件是 `zero/` 目录内 AI 协作规则的唯一真源，指导 Gin → go-zero 迁移与后续开发。

## 项目定位

本项目是 gin-vue-admin 后端的 go-zero 重写版本，目标是：

- 保持与原 Gin 后端 API 接口完全兼容（前端零改动）
- 使用 go-zero 框架规范，替换 Gin 全部基础设施
- 采用 goctl 生成脚手架，手写 logic 业务逻辑

## 架构总览

```
zero/
├── admin.go                 # 入口文件
├── admin.api                # API 定义（goctl DSL，路由 + 类型）
├── etc/
│   └── admin-api.yaml       # 配置文件
├── dao/                     # 数据访问层（GORM）
│   ├── model/               # GORM model 定义
│   │   ├── model.go         # 公共基础模型（GvaModel）
│   │   ├── sys_user.go      # 用户表
│   │   ├── sys_authority.go # 角色表
│   │   └── ...
│   └── repos/               # 数据库操作封装
│       ├── repos.go         # 基础 repo
│       ├── sys_user.go      # 用户 repo
│       └── ...
└── internal/
    ├── config/config.go     # 配置结构体
    ├── handler/             # HTTP 处理层（goctl 生成，一般不手动修改）
    │   ├── routes.go        # 路由注册（goctl 生成）
    │   ├── base/            # 按模块分组
    │   ├── user/
    │   └── ...
    ├── logic/               # 业务逻辑层（手写）
    │   ├── base/
    │   ├── user/
    │   └── ...
    ├── middleware/           # 中间件
    │   └── casbinhandlermiddleware.go
    ├── svc/
    │   └── servicecontext.go # 依赖注入容器
    └── types/
        └── types.go         # 请求/响应类型（goctl 生成）
```

## 分层规则

### handler 层

- 只做三件事：`Parse` → `Logic` → `Response`
- 禁止在 handler 中写业务逻辑
- 禁止在 handler 中直接访问数据库
- 由 goctl 自动生成，一般不手动修改
- 标准模式：

```go
func XxxHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.XxxReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        l := xxx.NewXxxLogic(r.Context(), svcCtx)
        resp, err := l.Xxx(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
        } else {
            httpx.OkJsonCtx(r.Context(), w, resp)
        }
    }
}
```

### logic 层

- 只写业务逻辑，不做 HTTP 参数解析
- 必须接收 `context.Context` 作为第一个参数
- 通过 `svcCtx` 访问配置、model、其他服务
- 函数超过 80 行必须拆分
- 标准模式：

```go
type XxxLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewXxxLogic(ctx context.Context, svcCtx *svc.ServiceContext) *XxxLogic {
    return &XxxLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *XxxLogic) Xxx(req *types.XxxReq) (resp *types.XxxResp, err error) {
    // 业务逻辑
    return
}
```

### model 层（dao/model）

- 存放 GORM model 定义，保持与原 Gin 项目数据库表结构一致
- 使用 `GvaModel` 基础模型（包含 ID、CreatedAt、UpdatedAt、DeletedAt）
- 每个 model 文件定义一个或相关的一组表结构
- 命名规则：表名用 `TableName()` 方法显式指定
- 新增表时参照原 `server/model/` 目录下的 GORM 定义

### repos 层（dao/repos）

- 封装数据库操作，每个 model 对应一个 repos 文件
- 通过 `*gorm.DB` 注入，支持事务和 context
- 只做数据访问，不包含业务逻辑
- 错误使用 `github.com/pkg/errors` 包装并补充上下文
- 标准模式：

```go
type XxxRepos struct {
    *BaseRepos
}

func NewXxxRepos(db *gorm.DB) *XxxRepos {
    return &XxxRepos{BaseRepos: NewBaseRepos(db)}
}

func (r *XxxRepos) FindById(ctx context.Context, id uint) (*model.Xxx, error) {
    var xxx model.Xxx
    err := r.DB.WithContext(ctx).Where("id = ?", id).First(&xxx).Error
    if err != nil {
        return nil, errors.Wrap(err, "查询xxx失败")
    }
    return &xxx, nil
}
```

### svc 层

- 依赖注入容器，持有所有共享依赖（Config、Middleware、Repos、GORM DB）
- 通过 `NewServiceContext` 初始化
- 新增 repos 或外部服务依赖时，在此注册

```go
type ServiceContext struct {
    Config        config.Config
    CasbinHandler rest.Middleware
    DB            *gorm.DB
    SysUserRepos  *repos.SysUserRepos
    // 新增 repos 在这里注册
}
```

### config 层

- 配置结构体在 `internal/config/config.go`
- 配置文件在 `etc/` 目录，YAML 格式
- 必须嵌入 `rest.RestConf`
- 验证配置：`run_mcp → validate_config`

## .api 文件规范

- `.api` 文件是 go-zero 的 DSL，定义路由和类型
- 由 goctl 生成，新增接口时应先修改 `.api` 文件再重新生成
- 类型定义使用 `type ( ... )` 块
- 路由定义使用 `@server` 指令声明分组、JWT、中间件、前缀
- `goctl api go -api admin.api -dir . -style go_zero` 重新生成代码

### @server 指令

```
@server(
    jwt:        Auth           # JWT 认证名，对应 config.Auth
    group:      moduleName     # 模块分组名
    prefix:     /modulePath    # 路由前缀
    middleware: CasbinHandler  # 中间件名，对应 svc 中的 middleware
)
```

### 新增接口流程

1. 在 `admin.api` 中添加类型定义和路由
2. 运行 `goctl api go -api admin.api -dir . -style go_zero`
3. 在 `internal/logic/xxx/` 中实现业务逻辑
4. handler 和 types 由 goctl 自动更新

## 错误处理

- 使用 `github.com/pkg/errors` 包装错误
- 错误只在边界层（logic）打印一次日志，其他层只 wrap
- wrap 使用 `errors.Wrap` / `errors.Wrapf`，必须补充上下文
- 禁止吞错，禁止 panic 处理业务错误
- 业务错误与系统错误区分：业务错误返回 `nil, errors.New("xxx")`，系统错误 wrap 后返回
- go-zero handler 层通过 `httpx.ErrorCtx` 自动处理 error 返回

## 日志规范

- 使用 go-zero `logx`
- 错误日志使用 `logx.Errorw`，不使用 `logx.Errorf` 字符串拼接
- 日志必须结构化，字段清晰

```go
logx.Errorw("login failed",
    logx.Field("error", err),
    logx.Field("trace_id", traceutil.TraceIDFromCtx(l.ctx)),
    logx.Field("module", "base"),
    logx.Field("action", "login"),
)
```

- 非边界层不随意打印 error 日志
- 常用字段：`error`、`trace_id`、`user_id`、`module`、`action`

## JSON 规范

- JSON 统一使用 `github.com/bytedance/sonic`
- 不使用 `encoding/json`，除非必须兼容标准库接口
- 序列化/反序列化错误必须返回并带上下文

## 代码质量

- `context.Context` 必须作为第一个参数
- 函数超过 80 行必须拆分
- 禁止重复代码
- 禁止 magic number / magic string
- 接口定义在使用方
- 不使用全局变量
- 必须符合 goimports 格式
- 必须通过 golangci-lint

## MCP-zero 工具使用

本项目优先使用 mcp-zero MCP 工具辅助开发：

| 工具 | 用途 | 时机 |
|------|------|------|
| `analyze_project` | 分析项目结构和依赖 | 每次大规模修改前 |
| `create_api_service` | 创建新的 API 服务 | 初始化新服务 |
| `create_api_spec` | 创建 API 规范文件 | 新增接口组 |
| `generate_api_from_spec` | 从规范生成代码 | 更新接口后 |
| `generate_model` | 从数据库表生成 model | 新增数据表 |
| `generate_config_template` | 生成配置模板 | 新增服务配置 |
| `generate_template` | 生成中间件/错误处理模板 | 新增通用组件 |
| `validate_config` | 验证配置文件 | 修改配置后 |
| `query_docs` | 查询 go-zero 文档 | 遇到框架问题时 |

## 外部参考

- [go-zero 官方文档](https://go-zero.dev)
- [ai-context](https://github.com/zeromicro/ai-context) — go-zero AI 上下文参考
- [zero-skills](https://github.com/zeromicro/zero-skills) — go-zero 技能集

## 统一响应格式

保持与原 Gin 后端一致：

```json
{
    "code": 0,
    "data": {},
    "msg": "success"
}
```

分页响应：

```json
{
    "code": 0,
    "data": {
        "list": [],
        "total": 0,
        "page": 1,
        "pageSize": 10
    },
    "msg": "success"
}
```

## 迁移注意事项

- 前端不改动，保持 API 路径和字段名完全一致
- 原 Gin 的 `gin.Context` 操作全部替换为 go-zero 的 `httpx` + `types`
- **model 层保持 GORM**，直接从 `server/model/` 复制并适配
- 原 `service/` 层替换为 `logic/` 层
- 原 `service/` 中的数据库操作替换为 `dao/repos/` 调用
- 原 `router/` 层替换为 `.api` DSL + `handler/routes.go`
- 原 `request/` 层替换为 `types/types.go`（goctl 生成）
- 原 `api/` 层替换为 `handler/`（goctl 生成）+ `logic/`（手写）
- Casbin 权限中间件需适配 go-zero 的 middleware 模式

## 文档维护

- 本文件只保留 go-zero 后端项目的高层、稳定规则
- 迁移过程中的具体映射关系写入 `aiDoc/` 对应目录
- 新增模块的示例代码放入 `aiDoc/examples/zero/`
