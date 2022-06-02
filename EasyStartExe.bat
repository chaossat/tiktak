go build main.go

cd service/favoriteaction
go build server.go

cd ../favoritelist
go build server.go

cd ../feed
go build server.go

cd ../followaction
go build server.go

cd ../followerlist
go build followerlistserver.go

cd ../followlist
go build server.go

cd ../login
go build loginServer.go

cd ../publist
go build publistService.go

cd ../register
go build server.go

cd ../userinf
go build userinfService.go

cd ../comment
go build server.go
cd ../..

start  main.exe
start  service/favoriteaction/server.exe
start  service/favoritelist/server.exe
start  service/feed/server.exe
start  service/followaction/server.exe
start  service/followerlist/followerlistserver.exe
start  service/followlist/server.exe
start  service/login/loginServer.exe
start  service/publist/publistService.exe
start  service/register/server.exe
start  service/userinf/userinfService.exe
start  service/comment/server.exe