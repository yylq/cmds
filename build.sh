export GO111MODULE=on
export GOPROXY=https://goproxy.cn
OUTPUT=$(pwd)/bin
cd cmds && go build -o ${OUTPUT}/cmds
