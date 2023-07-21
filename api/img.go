package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"os"
	"regexp"
	"time"
)

// @Summary 上传图片
// @Schemes
// @Description 可以上传多个图片
// @Tags img
// @Param img body any true "图片"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[any] 返回图片信息
// @Router /img/upload [post]
func noopImgUpload() {}

func imgRouter(app *gin.RouterGroup) {
	app.POST("/create", func(c *gin.Context) {
		//	图片数据接收,可以接收多个图片
		form, err := c.MultipartForm()
		if err != nil {
			curd.Error(c, errors.New("form接收失败"))
			return
		}
		file := form.File["file"]
		//	当前跟路径
		path, err := os.Getwd()
		if err != nil {
			curd.Error(c, errors.New("根路径获取失败"))
			return
		}
		t := time.Now()
		filePath := fmt.Sprintf("%s/static/%d/%d/", path, t.Year(), int(t.Month()))
		//存储文件
		resUrl := make(map[string]string)
		for i, f := range file {
			fileName, err := reFileName(f.Filename, i+1)
			if err != nil {
				curd.Error(c, err)
				return
			}
			resUrl[f.Filename] = filePath + fileName
			err = c.SaveUploadedFile(f, filePath+fileName)
			if err != nil {
				curd.Error(c, errors.New("图片存储失败"))
				return
			}
		}
		//	返回url: {传过来的文件名：/img/年/月/修改后的文件名}
		curd.OK(c, resUrl)
	})

}
func reFileName(f string, i int) (string, error) {
	format := `\.\w+$`
	re, err := regexp.Compile(format)
	if err != nil {
		return f, errors.New("图片格式正则没有匹配到")
	}
	index := re.FindStringIndex(f)
	if index == nil {
		return f, errors.New("没有找到图片格式的正则下标")
	}
	t := time.Now()
	fileName := fmt.Sprintf("%d%d%d%d%d", t.Day(), t.Hour(), t.Minute(), t.Second(), i)

	//f = fileName + f[index[0]:] //加了文件后缀
	f = fileName //没加了文件后缀
	return f, nil
}
