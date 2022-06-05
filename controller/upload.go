package controller

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/chaossat/tiktak/db"
	"github.com/chaossat/tiktak/middleware"
	"github.com/chaossat/tiktak/model"
	"github.com/chaossat/tiktak/oss"
	"github.com/chaossat/tiktak/util"
	"github.com/gin-gonic/gin"
)

//UploadHandler:处理视频上传请求
func UploadHandler(ctx *gin.Context) {
	//开启debug，观察性能瓶颈
	debugid, ok := <-DebugChan
	if ok {
		now := time.Now()
		log.Println("开始上传请求,操作ID:", debugid)
		defer log.Println("结束上传请求,操作ID:", debugid, "操作耗时：", time.Since(now))
	}

	//根据token鉴权，并获取userID
	token := ctx.PostForm("token")
	title := ctx.PostForm("title")
	user, err := middleware.CheckToken(token)
	if err != nil {
		fmt.Printf("Failed To Verify Token, err:%s\n", err.Error())
		UploadResponse(ctx, -1, "Invalid Token!")
		return
	}
	userID := user.UserID

	//获取文件
	file, header, err := ctx.Request.FormFile("data")
	if err != nil {
		fmt.Println(err.Error())
		UploadResponse(ctx, -3, "Error Occoured!")
		return
	}
	defer file.Close()
	if header.Size > 100<<20 {
		UploadResponse(ctx, 1, "文件不能大于100Mb!")
		return
	}

	//判断文件的后缀名，目前仅放行mp4文件
	ext := path.Ext(header.Filename)
	if ext != ".mp4" {
		UploadResponse(ctx, -4, "Only MP4 File Is Allowed!")
		return
	}

	//获取文件Sha1
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Printf("Failed to get file data, err:%s\n", err.Error())
		UploadResponse(ctx, -5, "Error Occoured!")
		return
	}
	sha1 := util.Sha1(buf.Bytes())

	//将文件写入临时存储位置
	tempLocation := "./tempfile/" + sha1 + ".mp4" // 临时存储地址
	newFile, err := os.Create(tempLocation)
	if err != nil {
		fmt.Printf("Failed to create file, err:%s\n", err.Error())
		UploadResponse(ctx, -6, "Error Occoured!")
		return
	}
	// defer newFile.Close()
	_, err = newFile.Write(buf.Bytes())
	if err != nil {
		fmt.Printf("Failed to copy file, err:%s\n", err.Error())
		UploadResponse(ctx, -7, "Error Occoured!")
		newFile.Close()
		return
	}
	//为视频生成封面
	util.CoverGenerator(tempLocation[2:], sha1+".jpg")
	cover, err := os.Open("./tempimage/" + sha1 + ".jpg")
	if err != nil {
		fmt.Printf("Failed to open cover file,maybe the video type is wrong, err:%s\n", err.Error())
		UploadResponse(ctx, -8, "Error Occoured! Maybe File Type Is Wrong")
		newFile.Close()
		os.Remove(tempLocation)
		return
	}
	//将视频信息存入数据库
	videoMeta := model.Video{
		Title:          title,
		AuthorID:       userID,
		UpdateTime:     time.Now().Unix(),
		Cover_location: "tempimage/" + sha1 + ".jpg",
		Location:       tempLocation[2:],
	}
	err = db.VideoUpload(&videoMeta)
	if err != nil {
		fmt.Printf("Failed to update mysql, err:%s\n", err.Error())
		UploadResponse(ctx, -9, "Error Occoured!")
		newFile.Close()
		return
	}

	//告诉客户端已经上传完成，并异步进行本地转存oss的操作
	UploadResponse(ctx, 0, "Upload Succeed!")

	//向通道压入转存请求
	ossPath := "videos/" + sha1 + ".mp4"
	ossImagePath := "images/" + sha1 + ".jpg"
	videoMeta.Location = ossPath
	videoMeta.Cover_location = ossImagePath
	videoOBJ := &oss.VideoOBJ{
		File:      newFile,
		Cover:     cover,
		VideoMeta: videoMeta,
	}
	oss.MQ_channel <- videoOBJ
}

//UploadResponse:返回上传处理信息
func UploadResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(200, gin.H{
		"status_code": code,
		"status_msg":  message,
	})
}
