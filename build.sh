#!/bin/bash
# windows
#gox -osarch="darwin/amd64" -output bin/feibor_darwin_amd64 ./main

work_path=$(cd `dirname $0`; pwd)

echo $work_path

cd $work_path

mkdir -p bin

# 编译Linux 和windows版本
#gox -osarch="darwin/amd64" -output ${work_path}/bin/operation_darwin_amd64 ./runner
#gox -osarch="windows/amd64" -output ${work_path}/bin/operation_win_amd64 ./runner
gox -osarch="linux/amd64" -output ${work_path}/proc ./runner

#GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./proc ./runner/proc.go

echo "done!"