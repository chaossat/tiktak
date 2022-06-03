#!/bin/bash

# go run main.go & \
# go run service/comment/server.go &\
# go run service/favoriteaction/server.go & \
# go run service/favoritelist/server.go & \
# go run service/feed/server.go & \
# go run service/followaction/server.go & \
# go run service/followerlist/followerlistserver.go & \
# go run service/followlist/server.go & \
# go run service/login/loginServer.go & \
# go run service/publist/publistService.go & \
# go run service/register/server.go & \
# go run service/userinf/userinfService.go & \
# echo "ok"

SHELL_PATH=$(cd "$(dirname "$0")";pwd)
# echo $SHELL_PATH
services=(comment favoriteaction favoritelist feed followaction followerlist followlist login publist register userinf)
# len=${#services[@]}
# echo $len
for  service in ${services[*]}
do 
    echo build:$service
    mkdir -p $SHELL_PATH/TiktakRelase/$service 
    cd $SHELL_PATH/service/$service/ &&\
    go build -o server &&\
    mv server $SHELL_PATH/TiktakRelase/$service/
done
cd $SHELL_PATH
go build -o main main.go 
mv main $SHELL_PATH/TiktakRelase/

echo  "用于记录PID">PID.txt
for  service in ${services[*]}
do 
    echo run:$service
    cd $SHELL_PATH
    ./TiktakRelase/$service/server -d=true >> PID.txt
done
cd $SHELL_PATH
./TiktakRelase/main -d=true >>PID.txt