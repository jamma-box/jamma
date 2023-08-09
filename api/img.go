package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
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

const imgRoot = "static/image"

type listParam struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`
}
type listRes struct {
	listParam
	Path string `json:"path"`
}

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
		filePath := fmt.Sprintf("%s/static/image/%d/%d/", path, t.Year(), int(t.Month()))
		//存储文件
		resUrl := make(map[string]string)
		for i, f := range file {
			fileName, err := ReFileName(f.Filename, i+1)
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
	app.GET("/list", func(c *gin.Context) {
		//绑定参数
		query := new(listParam)
		err := c.BindQuery(query)
		if err != nil {
			curd.Error(c, errors.New("参数绑定错误"))
			return
		}
		//	查询文件名
		res := make([]listRes, 0)
		filepath.WalkDir(imgRoot, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				arr := strings.Split(path, "\\")
				index := indexOf(arr, "image")
				res = append(res, listRes{
					listParam: listParam{
						Year:  arr[index+1],
						Month: arr[index+2],
						Day:   d.Name()[:2],
					},
					Path: path,
				})
			}
			return nil
		})
		//按查询参数返回字符串数组

		curd.List(c, res, int64(len(res)))
	})
}
func ReFileName(f string, i int) (string, error) {
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
	fileName := fmt.Sprintf("%v%d%d%d%d", t.Format("02"), t.Hour(), t.Minute(), t.Second(), i)

	//f = fileName + f[index[0]:] //加了文件后缀
	f = fileName //没加了文件后缀
	return f, nil
}
func indexOf(arr interface{}, target interface{}) int {
	slice := reflect.ValueOf(arr)
	for i := 0; i < slice.Len(); i++ {
		if slice.Index(i).Interface() == target {
			return i
		}
	}
	return -1
}
