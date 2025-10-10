# 临时设置环境变量（仅当前终端有效）
export GOOS=linux   # 目标操作系统为 Linux
export GOARCH=amd64 # 目标架构为 64位 x86（Ubuntu 服务器常用）

# 编译你的程序（例如编译 main.go）
go build


