export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOCACHE=$(pwd)/.cache
OUTPUT=$(pwd)/bin
cd cmds && go build -o ${OUTPUT}/cmds
