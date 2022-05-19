package publist

import (
	"context"

	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/oss"
	"github.com/chaossat/tiktak/service/publist/pb"
)

type PublistServer struct {
	pb.UnimplementedPublishServer
}

// 获取发布的video列表
func (this *PublistServer) Publist(ctx context.Context, req *pb.DouyinPublishListRequest) (*pb.DouyinPublishListResponse, error) {
	// token验证
	_, err := middleware.CheckToken(*req.Token)
	if err != nil {
		var code int32 = -1
		var msg string = "验证失败！"
		return &pb.DouyinPublishListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, nil
	}
	userinf, err := db.UserInfoById(int(*req.UserId))
	//TODO:消费err
	videoes, err := db.VideoedByID(int(*req.UserId))
	//TODO:消费err
	isfollow := false
	var user = pb.User{
		Id:            &userinf.ID,
		Name:          &userinf.Username,
		FollowCount:   &userinf.Follow_count,
		FollowerCount: &userinf.Follower_count,
		IsFollow:      &isfollow,
	}
	var vide_ans = make([]*pb.Video, len(videoes))
	for i := 0; i < len(videoes); i++ {
		var isfabvorite = false
		playURL := oss.GetURL(videoes[i].Location)
		coverURL := oss.GetURL(videoes[i].Cover_location)
		vide_ans[i] = &pb.Video{
			Id:            &videoes[i].ID,
			Author:        &user,
			PlayUrl:       &playURL,
			CoverUrl:      &coverURL,
			FavoriteCount: &videoes[i].FavoriteCount,
			CommentCount:  &videoes[i].CommentCount,
			IsFavorite:    &isfabvorite,
		}
	}
	if err != nil {
		var code int32 = -2
		var msg string = "视频查找失败！"
		return &pb.DouyinPublishListResponse{
			StatusCode: &code,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, nil
	}
	var code int32 = 0
	var msg string = "验证成功！"
	res := &pb.DouyinPublishListResponse{
		StatusCode: &code,
		StatusMsg:  &msg,
		VideoList:  vide_ans,
	}
	return res, nil
}
