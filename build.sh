export GOPROXY=https://goproxy.cn
BASE=$(pwd)
echo $BASE
cd cmds && go build -o ../bin/cmds
