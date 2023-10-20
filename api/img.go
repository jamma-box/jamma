package api

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
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
		getwd, err := os.Getwd()
		if err != nil {
			curd.Error(c, errors.New("根路径获取失败"))
			return
		}
		t := time.Now()
		filePath := filepath.Join(getwd, "static", "image", fmt.Sprintf("%d", t.Year()), fmt.Sprintf("%d", t.Month()))
		//存储文件
		resUrl := make([]string, 0)
		for _, f := range file {
			//时间戳命名
			filename := fmt.Sprintf("%v.png", time.Now().UnixMilli())
			//响应添加
			resUrl = append(resUrl, filepath.Join(strings.TrimPrefix(filePath, getwd), filename))
			//转存
			err = c.SaveUploadedFile(f, filepath.Join(filePath, filename))
			if err != nil {
				curd.Error(c, errors.New("图片存储失败:"+err.Error()))
				return
			}
		}
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
		t := time.Now()
		if query.Year == "" {
			query.Year = fmt.Sprintf("%v", t.Year())
		}
		if query.Month == "" {
			query.Month = fmt.Sprintf("%v", int(t.Month()))
		}
		if query.Day == "" {
			query.Day = fmt.Sprintf("%v", t.Day())
		}
		//	查询文件名
		getwd, err := os.Getwd()
		if err != nil {
			curd.Error(c, err)
			return
		}
		imgRoot := filepath.Join(getwd, "static", "image")
		res := make([]listRes, 0)
		err = filepath.WalkDir(imgRoot, func(fp string, fi fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !fi.IsDir() {
				filename := filepath.Base(fp)
				ext := filepath.Ext(filename)
				name := strings.TrimSuffix(filename, ext)
				timeu, err := strconv.ParseInt(name, 10, 64)
				if err != nil {
					return err
				}
				t := time.UnixMilli(timeu)
				res = append(res, listRes{
					listParam: listParam{
						Year:  fmt.Sprintf("%d", t.Year()),
						Month: fmt.Sprintf("%d", t.Month()),
						Day:   fmt.Sprintf("%d", t.Day()),
					},
					Path: strings.TrimPrefix(fp, getwd),
				})
			}
			return nil
		})
		if err != nil {
			curd.Error(c, err)
			return
		}
		curd.List(c, res, int64(len(res)))
	})
}
