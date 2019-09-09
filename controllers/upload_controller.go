package controllers

import (
	"blog.api/models/dto/response"
	"os"
	"strings"
	"time"
	"fmt"
	"path/filepath"
	"io"

)

//文件上传控制器
type UploadController struct {
	BaseController
}

// @Title Post
// @Description 文件上传
// @Param	fileUpload	 form 	string	true		"文件上传的数据"
// @Success 200 {object} response.FileUploadResult
// @Failure 403 :id is empty
// @router / [post]
func (this *UploadController) Post() {

	fmt.Println("fileupload...")
	fileData, fileHeader, err := this.GetFile("fileUpload")
	defer fileData.Close()
	if err != nil {
		this.responseErr(err)
		return
	}

	now := time.Now()

	fileType := "other"
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())

	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		this.responseErr(err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		this.responseErr(err)
		return
	}

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		this.responseErr(err)
		return
	}

	if fileType == "img" {

	}

	var result response.FileUploadResult
	result.Code=200
	result.Message ="上传成功"
	result.Url=filePathStr

	this.Data["json"] = result
	this.ServeJSON()

}

func (this *UploadController) responseErr(err error) {
	this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	this.ServeJSON()
}


