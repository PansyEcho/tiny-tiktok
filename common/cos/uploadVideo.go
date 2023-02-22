package cos

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"tiny-tiktok/common/errx"
)

type UploaderVideo struct {
	UserId      int64
	MachineId   uint16
	VideoBucket string
	SecretID    string
	SecretKey   string
}

func (l *UploaderVideo) UploadVideo(ctx context.Context, file multipart.File) (string, error) {
	u, _ := url.Parse(l.VideoBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  l.SecretID,
			SecretKey: l.SecretKey,
		},
	})

	genSnowFlake := new(GenSnowFlake)
	id, err := genSnowFlake.GenSnowFlake(l.MachineId)
	if err != nil {
		logx.Errorf("UploadVideo--->GenSnowFlake err : %v", err)
		return "", err
	}
	// 生成useId/id.mp4
	key := strconv.FormatInt(l.UserId, 10) + "/" + strconv.FormatInt(int64(id), 10)
	// 上传视频文件
	_, err = c.Object.Put(ctx, key+".mp4", file, nil)
	if err != nil {
		logx.Errorf("UploadVideo--->Put err : %v", err)
		return "", err
	}

	// 上传成功 返回key
	return key, nil
}

func (l *UploaderVideo) SaveUploadedFile(ctx context.Context, file multipart.File, videoFileName string) (saveVideoPath, savePhotoPath string, err error) {

	saveVideoPath = fmt.Sprintf("%s%s", "/video/", videoFileName)
	savePhotoPath = fmt.Sprintf("%s%s.jpg", "/photo/", strings.Split(videoFileName, ".")[0])
	// 保存视频到临时文件夹
	tmpVideoPath := fmt.Sprintf("%s/dousheng-%s", os.TempDir(), videoFileName)
	// perm权限
	fileByte, err := fileToBytes(file)
	if err != nil {
		return "", "", err
	}
	err = os.WriteFile(tmpVideoPath, fileByte, 0666)
	if err != nil {
		err = fmt.Errorf("上传失败：%w", err)
		return
	}
	defer os.Remove(tmpVideoPath)
	// TODO 队列防ffmpeg并发冲突
	// 使用 cmd 命令调用 ffmpeg 生成截图 ，传入的参数一为 视频的真实路径，参数二为生成图片保存的真实路径
	tempPhotoPath := fmt.Sprintf("%s/test-%s.jpg", os.TempDir(), strings.Split(videoFileName, ".")[0])
	cmd := exec.Command("ffmpeg", "-i", tmpVideoPath, tempPhotoPath,
		"-ss", "00:00:00", "-r", "1", "-vframes", "1", "-an", "-vcodec", "mjpeg")
	err = cmd.Run()
	if err != nil {
		fmt.Println("ffmpeg报错")
		return "", "", err
	}
	defer os.Remove(tempPhotoPath)

	u, _ := url.Parse(l.VideoBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  l.SecretID,
			SecretKey: l.SecretKey,
		},
	})

	genSnowFlake := new(GenSnowFlake)
	id, err := genSnowFlake.GenSnowFlake(l.MachineId)
	if err != nil {
		logx.Errorf("UploadVideo--->GenSnowFlake err : %v", err)
		return "", "", err
	}
	//生成useId/id.mp4
	key := strconv.FormatInt(l.UserId, 10) + "/" + strconv.FormatInt(int64(id), 10)
	// 上传视频文件
	fmt.Println("开始上传")
	_, err = c.Object.Put(ctx, key+".mp4", file, nil)
	if err != nil {
		logx.Errorf(errx.MapErrMsg(errx.UPLOAD_VIDEO_ERROR), err)
		return "", "", err
	}
	fmt.Println("上传成功")

	//上传截图
	//_, err = c.Object.PutFromFile(ctx, savePhotoPath, tempPhotoPath, nil)
	//if err != nil {
	//	logx.Errorf(errx.MapErrMsg(errx.UPLOAD_VIDEO_ERROR), err)
	//	return "", "", err
	//}

	return
}

func fileToBytes(file multipart.File) ([]byte, error) {
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
