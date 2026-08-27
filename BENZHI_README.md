# 多语种人文语料生命周期服务

这是一个面向外语文学、语言学与计算机交叉研究的 Go 后端。服务管理语料、许可、标注批次、分析记录和审计事件。

## 构建与运行

需要 Go 1.26 或更新版本。执行 `go mod download` 后运行 `go run .`，默认监听 `:8080`。可通过 `ADDR`、`DATABASE_PATH` 和 `WORKER_INTERVAL_SECONDS` 配置。

## 测试

执行 `go test ./... -count=1`、`go test -race ./... -count=1`、`go vet ./...` 和 `go build ./...`。

## Docker

执行 `docker build -t digital-humanities-go .` 构建镜像，容器默认入口运行服务。`/healthz` 提供存活检查，`/readyz` 提供数据库就绪检查。
