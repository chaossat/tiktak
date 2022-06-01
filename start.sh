#!/bin/bash
go run main.go & \
go run service/favoriteaction/server.go & \
go run service/favoritelist/server.go & \
go run service/feed/server.go & \
go run service/followaction/server.go & \
go run service/followerlist/followerlistserver.go & \
go run service/followlist/server.go & \
go run service/login/loginServer.go & \
go run service/publist/publistService.go & \
go run service/register/server.go & \
go run service/userinf/userinfService.go & \