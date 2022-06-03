#!/bin/bash
SHELL_PATH=$(cd "$(dirname "$0")";pwd)
services=(comment favoriteaction favoritelist feed followaction followerlist followlist login publist register userinf)
echo  "用于记录PID">PID.txt
for  service in ${services[*]}
do 
    echo run:$service
    cd $SHELL_PATH
    ./TiktakRelase/$service/server -d=true >> PID.txt
done
cd $SHELL_PATH
./TiktakRelase/main -d=true >>PID.txt