# ginapp
一个 gin 的 golang 项目实例

## 使用方法
```sh
cd ~/projects
mkdir ginapp
cd ginapp
# 初始化 golang 项目
go mod init github.com/zoujiaqing/ginapp
# 引入 gin 依赖
go get -u github.com/gin-gonic/gin
# 安装依赖
go install
# 编译项目
go build -o server main.go
# 运行项目
./server
```