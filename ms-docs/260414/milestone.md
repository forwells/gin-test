# 里程碑 01 — Gin 项目基础结构搭建

**日期：** 2026-04-14  
**项目：** gin-test（小型博客系统训练项目）  
**背景：** PHP Laravel 主项目 + Go 微服务模块的混合架构学习路线

---

## 当日完成内容

### 1. 项目结构重组

将原本全部堆在 `main.go` 里的代码拆分为分层结构：

```
gin-test/
├── main.go           # 只负责启动，调用 router.Setup()
├── go.mod
├── .air.toml         # air 热重载配置
└── router/
    ├── router.go     # 路由汇总入口，Setup() 函数
    ├── test.go       # 测试路由（/test）
    └── posts.go      # 文章路由（/posts CRUD，handler 占位 nil）
```

### 2. 当前各文件内容

**`main.go`**
```go
package main

import "gin-test/router"

func main() {
    r := router.Setup()
    r.Run(":6063")
}
```

**`router/router.go`**
```go
func Setup() *gin.Engine {
    r := gin.Default()
    regTestRoutes(r)
    return r
}
```

**`router/test.go`**
```go
func regTestRoutes(r *gin.Engine) {
    test := r.Group("/test")
    test.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"msg": "你好 gin!"})
    })
}
```

**`router/posts.go`**
```go
func regPostRoutes(r *gin.Engine) {
    posts := r.Group("/posts")
    posts.GET("/", nil)
    posts.GET("/:id", nil)
    posts.POST("/", nil)
    posts.PUT("/:id", nil)
    posts.DELETE("/:id", nil)
}
```

---

## 当日掌握知识点

### Go 语言基础
- **值接收者 vs 指针接收者**
  - `func (c XXStruct) F()` → 操作副本，原始数据不变
  - `func (c *XXStruct) F()` → 操作原始数据，Go 自动处理取址
  - 实际项目中统一用指针接收者，Gin 的 `Context` 就是典型案例
- **package 命名规则**
  - 同一目录下所有 `.go` 文件必须使用同一个 package 名
  - package 名与目录名保持一致（约定，非强制，但务必遵守）
  - `main` 包是唯一例外，不能被其他包 import
- **import 路径规则**
  - 路径来自 `go.mod` 的 `module` 名作为根
  - `gin-test/router` → 对应 `router/` 目录
  - import 写目录路径，调用时用 package 名
- **可见性规则**
  - 大写开头 → exported，包外可调用
  - 小写开头 → unexported，仅包内使用
  - `regPostRoutes` 小写：只在 router 包内被调用，不对外暴露

### Gin 框架
- **`c.JSON()` 调用链**：`c.JSON()` → `c.Render()` → `render.JSON{Data}` → `encoding/json.Marshal` → 写入响应体
- **`gin.H`** 本质是 `map[string]any` 的类型别名
- **路由分组** `r.Group("/prefix")` 等价于给该组所有路由加公共前缀
- **`gin.Default()`** 包含 Logger 和 Recovery 两个默认中间件

### 工具链
- **`air`** 热重载工具：监听 `.go` 文件变化，自动重新编译并重启
  - 安装：`go install github.com/air-verse/air@latest`
  - 需要将 `$GOPATH/bin` 加入 `$PATH`：`echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc`
  - 项目根目录执行 `air init` 生成 `.air.toml` 配置文件

### 项目结构设计
- **`internal/` 目录**：Go 编译器级别的访问控制，外部包无法 import，核心业务逻辑放这里
- **`pkg/` 目录**：可复用工具包，可被外部 import
- 路由按模块拆分文件（`posts.go` / `user.go` / `comment.go`），`router.go` 只做汇总

---

## 遗留问题 / 下一步

- [ ] `router/router.go` 的 `Setup()` 还未注册 `regPostRoutes`（posts 路由未挂载）
- [ ] handler 层：路由处理函数目前都是 `nil` 占位，需要抽离到 `internal/handler/`
- [ ] 统一响应格式封装（`pkg/response/`）：`code` / `message` / `data`
- [ ] 请求参数绑定与验证（对标 Laravel FormRequest）
- [ ] 连接 PostgreSQL（GORM）
- [ ] 用户认证（JWT）

---

## 学习路线进度

| 阶段 | 内容 | 状态 |
|------|------|------|
| 第一阶段 | Gin 基础：路由、中间件、项目结构 | 🔄 进行中 |
| 第二阶段 | 数据库 GORM + 配置管理 viper | ⏳ 待开始 |
| 第三阶段 | OpenAPI 规范 + Laravel 对接 | ⏳ 待开始 |
| 第四阶段 | 高并发 / 队列 / WebSocket | ⏳ 待开始 |
| 第五阶段 | Docker 部署 | ⏳ 待开始 |
