package controllers

import (
	"path"
	"time"
	"viv/models"
	"viv/vutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// FileController : 文件控制器
type FileController struct {
	beego.Controller
}

// UploadFiles : 多图上传
func (c *FileController) UploadFiles() {

	fileHeads, err := c.GetFiles("images")
	if err != nil {
		beego.Info(err)
		return
	}
	beego.Info(fileHeads)
	o := orm.NewOrm()
	for _, image := range fileHeads {
		imageExt := path.Ext(image.Filename)
		if imageExt != ".png" && imageExt != ".jpg" && imageExt != ".jpeg" {
			beego.Info("文件格式不对")
			continue
		}

		if image.Size > 5*1024*1024 {
			beego.Info("file too large")
			continue
		}

		//时间转换到秒一下级别根据后面0的个数
		timeStr := time.Now().Format("2006-01-02 03:04:05.000000")
		fname := timeStr + imageExt
		imageURL := "/static/img/" + fname
		c.SaveToFile("images", "."+imageURL)
		err = insertFileToDBWithHash(o, "."+imageURL)
		// imageModel := models.Image{ImageStr: imageURL}
		// _, err = o.Insert(&imageModel)
		if err != nil {
			beego.Info(err)
			continue
		}
	}

	c.Ctx.WriteString("1234")
}

// insertFileToDBWithHash : 文件插入数据库 hash唯一 节省磁盘空间
func insertFileToDBWithHash(o orm.Ormer, filepath string) error {
	hashvalue, err := vutil.VHashValue(filepath)
	if err != nil {
		beego.Info(err)
		return err
	}
	imageModel := models.Image{Hash: hashvalue}
	err = o.Read(&imageModel, "Hash")
	if err == nil {
		beego.Info(err)
		return nil
	}
	imageModel.ImageStr = filepath
	_, err = o.Insert(&imageModel)
	return err
}

// UploadFile : 大文件上传
func (c *FileController) UploadFile() {
	f, h, err := c.GetFile("file")
	if err != nil {
		beego.Info(err)
		return
	}
	defer f.Close()
	fileExt := path.Ext(h.Filename)

	if fileExt != ".zip" && fileExt != ".rar" {
		beego.Info("file format error")
		return
	}

	timeStr := time.Now().Format("2006-01-02 03:04:05.000") + fileExt
	fileName := "./static/img/" + timeStr
	c.SaveToFile("file", fileName)
}
