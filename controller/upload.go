package controller

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/model"
	"github.com/chaossat/tiktak/oss"
	"github.com/chaossat/tiktak/util"
	"github.com/gin-gonic/gin"
)

//UploadHandler:处理视频上传请求
func UploadHandler(ctx *gin.Context) {
	//TODO:根据token鉴权，并获取userID
	userID := 1

	//获取文件
	file, header, err := ctx.Request.FormFile("data")
	if err != nil {
		fmt.Println(err.Error())
		UploadResponse(ctx, -1, "Error Occoured!")
		return
	}
	defer file.Close()

	//判断文件的后缀名，目前仅放行mp4文件
	ext := path.Ext(header.Filename)
	if ext != ".mp4" {
		UploadResponse(ctx, -2, "Only MP4 File Is Allowed!")
		return
	}

	//获取文件Sha1
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Printf("Failed to get file data, err:%s\n", err.Error())
		UploadResponse(ctx, -3, "Error Occoured!")
		return
	}
	sha1 := util.Sha1(buf.Bytes())

	//将文件写入临时存储位置
	tempLocation := "./tempfile/" + sha1 // 临时存储地址
	newFile, err := os.Create(tempLocation)
	if err != nil {
		fmt.Printf("Failed to create file, err:%s\n", err.Error())
		UploadResponse(ctx, -4, "Error Occoured!")
		return
	}
	defer newFile.Close()
	_, err = newFile.Write(buf.Bytes())
	if err != nil {
		fmt.Printf("Failed to copy file, err:%s\n", err.Error())
		UploadResponse(ctx, -5, "Error Occoured!")
		return
	}
	videoMeta := model.Video{
		Title:          header.Filename[:len(header.Filename)-4],
		AuthorID:       userID,
		FavorateCounts: 0,
		CommentCounts:  0,
		UpdateTime:     int(time.Now().Unix()),
		Location:       tempLocation[2:],
	}
	err = db.VideoUpload(&videoMeta)
	if err != nil {
		fmt.Printf("Failed to update mysql, err:%s\n", err.Error())
		UploadResponse(ctx, -6, "Error Occoured!")
		return
	}

	//告诉客户端已经上传完成，并继续进行上传oss的操作
	UploadResponse(ctx, 0, "Upload Succeed!")

	//继续上传oss,并更新文件Location
	//TODO：做成异步处理,然后在更新完成视频地址后，删除本地文件
	ossPath := "videos/" + sha1
	newFile.Seek(0, 0) // 游标重新回到newFile文件头部，否则读不出任何数据
	err = oss.Bucket().PutObject(ossPath, newFile)
	if err != nil {
		fmt.Println(newFile)
		fmt.Printf("Failed while pushing to oss, err:%s\n", err.Error())
		return
	}
	videoMeta.Location = "oss:" + ossPath
	err = db.VideoLocationUpdate(&videoMeta)
	if err != nil {
		fmt.Printf("Failed to update mysql, err:%s\n", err.Error())
	}
}

// func Upload_Oss(){
//
// }

//UploadResponse:返回上传处理信息
func UploadResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(200, gin.H{
		"status_code": code,
		"status_msg":  message,
	})
}
