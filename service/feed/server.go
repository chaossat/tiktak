package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/chaossat/tiktak/oss"
	"github.com/chaossat/tiktak/service/feed/model"
	"github.com/chaossat/tiktak/service/feed/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Feed struct {
}

//本次返回的视频中发布最早的时间，作为下次请求时的latest_time
func GetNextTime(videos []model.Video) int64 {
	if len(videos) == 0 {
		return time.Now().Unix()
	}
	ans := videos[0].UpdateTime
	for _, video := range videos {
		if ans < video.UpdateTime {
			ans = video.UpdateTime
		}
	}
	return ans
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

func (this *Feed) GetFeed(ctx context.Context, req *pb.DouyinFeedRequest) (*pb.DouyinFeedResponse, error) {
	var statuscode int32
	var statusmsg string
	var userid int64
	// if len(*req.Token) > 0 {
	// 	claims, err := middleware.CheckToken(*req.Token)
	// 	if err != nil {
	// 		log.Println("解析用户token失败", err.Error())
	// 		statuscode = 1
	// 		statusmsg = "解析用户token失败"
	// 		return &pb.DouyinFeedResponse{
	// 			StatusCode: &statuscode,
	// 			StatusMsg:  &statusmsg,
	// 		}, nil
	// 	}
	// 	userid = claims.UserID
	// } else {
	// 	userid = 0
	// }

	video_list, err := model.GetVideoList(*req.LatestTime)
	if err != nil {
		statuscode = 1
		statusmsg = "获取视频列表失败"
		return &pb.DouyinFeedResponse{
			StatusCode: &statuscode,
			StatusMsg:  &statusmsg,
		}, nil
	}
	var pbvideo_list []*pb.Video
	for _, video := range video_list {
		Id := video.ID
		Title := video.Title
		Author := GetAuthor(userid, video.AuthorID)
		PlayUrl := oss.GetURL(video.Location)
		CoverUrl := oss.GetURL(video.Cover_location)
		//favoritecnt := int64(0)
		favoritecnt, err := model.GetFavoriteCount(Id)
		if err != nil {
			log.Println("获取视频的点赞个数错误", err.Error())
			statuscode = 1
			statusmsg = "获取视频点赞失败"
			return &pb.DouyinFeedResponse{
				StatusCode: &statuscode,
				StatusMsg:  &statusmsg,
			}, nil
		}
		FavoriteCount := favoritecnt
		//commentcnt := int64(0)
		commentcnt := model.GetCommentCount(video)
		CommentCount := commentcnt
		var isfavorite bool
		if userid != 0 {
			isfavorite, err = model.IsFavorite(userid, video.ID)
			if err != nil {
				log.Println("判断用户是否点赞当前视频错误")
				statuscode = 1
				statusmsg = "判断用户是否点赞当前视频错误"
				return &pb.DouyinFeedResponse{
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
	statusmsg = "获取视频流成功"
	nexttime := GetNextTime(video_list)
	return &pb.DouyinFeedResponse{
		StatusCode: &statuscode,
		StatusMsg:  &statusmsg,
		VideoList:  pbvideo_list,
		NextTime:   &nexttime,
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
	log.Println("正在启动Feed服务......")
	InitConfig()
	model.InitDB()
	model.InitRedis()
	//初始化grpc实例
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterFeedServer(grpcServer, new(Feed))

	//设置监听
	listen, err := net.Listen("tcp", ":12346")
	if err != nil {
		log.Println("注册服务启动监听失败")
	}
	defer listen.Close()

	//启动服务
	grpcServer.Serve(listen)
}
