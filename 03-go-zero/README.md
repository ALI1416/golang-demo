1. 安装go-zero脚手架goctl `go install github.com/zeromicro/go-zero/tools/goctl`
2. 初始化项目test `go mod init test`
3. 安装go-zero框架 `go get -u github.com/zeromicro/go-zero`
4. 使用goctl生成服务demo `goctl api new demo`
5. 安装依赖 `go mod tidy`
6. 切到demo目录 `cd demo`
7. 修改文件`internal/logic/demologic.go`，新增返回
8. 运行服务 `go run demo.go`
9. 访问链接`http://localhost:8888/from/me`
10. 修改`demo.api`文件，新增`list/get`API
11. 重新生成服务`goctl api go --dir ./ --api demo.api`
12. 修改文件`internal/logic/listlogic.go`，新增返回
13. 运行服务 `go run demo.go`
14. 访问链接`http://localhost:8888/list/get`
