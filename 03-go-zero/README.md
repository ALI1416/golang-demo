1. 安装go-zero脚手架goctl `go install github.com/zeromicro/go-zero/tools/goctl`
2. 初始化项目test `go mod init test`
3. 安装go-zero框架 `go get -u github.com/zeromicro/go-zero`
4. 使用goctl生成服务demo `goctl api new demo`
5. 安装依赖 `go mod tidy`
6. 切到demo目录，运行服务 `go run demo.go`
