package model

import (
	"fmt"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func LinkqiniuSDK() {
	// 七牛云存储中我的 Access Key 和 Secret Key
	accessKey := "HgEHmfBLVUOU2ULPn9VU5ZQSJu9KeekVjonltQIF"
	secretKey := "6fE3itlwcX_P3q5sDXnzs8YyYbkHynUfXd963oyb"
	// 创建一个 MAC（管理凭证），用于签名
	mac := qbox.NewMac(accessKey, secretKey)
	// 要上传的空间名（Bucket 名称）
	bucket := "myTikTok"
	// 存储区域（Zone）
	zone := &storage.ZoneHuanan
	cfg := storage.Config{
		Zone:          zone, // 根据实际情况选择 Zone
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	// 创建一个七牛云存储对象
	formUploader := storage.NewFormUploader(&cfg)
	putExtra := storage.PutExtra{}

	// 本地文件路径
	localFile := "/path/to/local/file.txt"

	// 上传文件的唯一 key
	key := "file.txt"

	// 生成上传凭证
	// token := ""
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	// 上传文件
	err := formUploader.PutFile(nil, &storage.PutRet{}, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println("上传文件失败:", err)
		return
	}

	fmt.Println("上传文件成功!")
}
