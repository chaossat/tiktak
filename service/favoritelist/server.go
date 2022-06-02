package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/oss"
	"github.com/chaossat/tiktak/service/favoritelist/model"
	"github.com/chaossat/tiktak/service/favoritelist/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type FavoriteList struct {
}

//获取作者相关信息
func GetAuthor(userid, authorid int64) pb.User {
	var user model.User
	var pbauthor pb.User
	if userid != 0 {
		var err error
		user, err = model.GetUser(userid)
		if err != nil {
			log.Println("查询当前用户错误")
			return pbauthor
		}
	}
	author, err := model.GetUser(authorid)
	if err != nil {
		log.Println("查询视频作者错误")
		return pbauthor
	}
	Id := author.ID
	Name := author.Username
	followcount := model.GetFollowCount(author)
	FollowCount := followcount
	followercount := model.GetFollowerCount(author)
	FollowerCount := followercount
	var isfollow bool
	if userid == 0 {
		isfollow = false
	} else {
		isfollow, err = model.IsFollow(user, author)
		if err != nil {
			isfollow = false
		}
	}
	IsFollow := isfollow
	pbauthor = pb.User{
		Id:            &Id,
		Name:          &Name,
		FollowCount:   &FollowCount,
		FollowerCount: &FollowerCount,
		IsFollow:      &IsFollow,
	}
	return pbauthor
}

func (this *FavoriteList) GetFavoriteList(ctx context.Context, req *pb.DouyinFavoriteListRequest) (*pb.DouyinFavoriteListResponse, error) {
	var statuscode int32
	var statusmsg string
	user_id := *req.UserId
	video_list, err := model.FavoriteList(user_id)
	var curid int64
	if len(*req.Token) > 0 {
		claims, err := middleware.CheckToken(*req.Token)
		if err != nil {
			log.Println("解析用户token失败", err.Error())
			statuscode = 1
			statusmsg = "解析用户token失败"
			return &pb.DouyinFavoriteListResponse{
				StatusCode: &statuscode,
				StatusMsg:  &statusmsg,
			}, nil
		}
		curid = claims.UserID
	}
	//else {
	//	statuscode = 1
	//	statusmsg = "用户未登录"
	//	return &pb.DouyinFavoriteListResponse{
	//		StatusCode: &statuscode,
	//		StatusMsg:  &statusmsg,
	//	}, nil
	//}
	if err != nil {
		statuscode = 1
		statusmsg = "获取点赞过的视频列表失败"
		return &pb.DouyinFavoriteListResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
		}, nil
	}

	var pbvideo_list []*pb.Video
	for _, video := range video_list {
		Id := video.ID
		Title := video.Title
		Author := GetAuthor(curid, video.AuthorID)
		PlayUrl := oss.GetURL(video.PlayLocation)
		CoverUrl := oss.GetURL(video.Cover_location)
		favoritecnt, err := model.GetFavoriteCount(Id)
		if err != nil {
			log.Println("获取视频的点赞个数错误", err.Error())
			statuscode = 1
			statusmsg = "获取视频点赞失败"
			return &pb.DouyinFavoriteListResponse{
				StatusCode: &statuscode,
				StatusMsg:  &statusmsg,
			}, nil
		}
		FavoriteCount := favoritecnt
		commentcnt := model.GetCommentCount(video)
		CommentCount := commentcnt
		var isfavorite bool
		if curid != 0 {
			isfavorite, err = model.IsFavorite(curid, video.ID)
			if err != nil {
				log.Println("判断用户是否点赞当前视频错误")
				statuscode = 1
				statusmsg = "判断用户是否点赞当前视频错误"
				return &pb.DouyinFavoriteListResponse{
					StatusCode: &statuscode,
					StatusMsg:  &statusmsg,
				}, nil
			}
		} else {
			isfavorite = false
		}

		IsFavorite := isfavorite
		pbvideo := pb.Video{
			Id:            &Id,
			Title:         &Title,
			Author:        &Author,
			PlayUrl:       &PlayUrl,
			CoverUrl:      &CoverUrl,
			FavoriteCount: &FavoriteCount,
			CommentCount:  &CommentCount,
			IsFavorite:    &IsFavorite,
		}
		pbvideo_list = append(pbvideo_list, &pbvideo)
	}
	statuscode = 0
	statusmsg = "获取点赞过的视频成功"
	return &pb.DouyinFavoriteListResponse{
		StatusCode: &statuscode,
		StatusMsg:  &statusmsg,
		VideoList:  pbvideo_list,
	}, nil
}

func InitConfig() {
	workDir, _ := os.Getwd() //获取当前工作路径，非文件路径，以终端显示路径为准
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("正在启动FavoriteList服务......")
	InitConfig()
	model.InitDB()
	model.InitRedis()
	//初始化grpc实例
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterFavoriteListServer(grpcServer, new(FavoriteList))

	//设置监听
	listen, err := net.Listen("tcp", ":12350")
	if err != nil {
		log.Println("注册服务启动监听失败")
	}
	defer listen.Close()

	//启动服务
	grpcServer.Serve(listen)
}
