package utils

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bysir-zl/bygo/util/encoder"
	"github.com/tealeg/xlsx"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"golang_utils/glog"
)

func ObjToMap(obj interface{}) map[string]interface{} {
	pointer := reflect.Indirect(reflect.ValueOf(obj))
	typer := pointer.Type()

	fieldNum := pointer.NumField()

	data := map[string]interface{}{}

	for i := 0; i < fieldNum; i++ {
		field := pointer.Field(i)
		key := typer.Field(i).Name
		data[key] = field.Interface()
	}

	return data
}

func ObjListToMapList(obj interface{}) (mappers []map[string]interface{}) {
	mappers = []map[string]interface{}{}

	value := reflect.ValueOf(obj)
	for i := 0; i < value.Len(); i = i + 1 {
		item := value.Index(i)
		mappers = append(mappers, ObjToMap(item.Interface()))
	}
	return
}

func ParseHashStr(str string) string {
	hvBytes := []byte(str)
	nBytes := []byte("")
	level := 3
	for i, v := range hvBytes {
		if level > 0 && i%2 == 0 {
			nBytes = append(nBytes, '/')
			level--
		}
		nBytes = append(nBytes, v)
	}
	return string(nBytes)
}

func ServerExcelStruct(data interface{}, name string, col []string, maps map[string]string, c beego.Controller) (error) {
	return ServerExcelMap(ObjListToMapList(data), name, col, maps, c)
}
func ServerExcelMap(data []map[string]interface{}, name string, col []string, maps map[string]string, c beego.Controller) (err error) {
	file := xlsx.NewFile()
	sht, err := file.AddSheet(name)
	if err != nil {
		return
	}
	if col == nil {
		if data == nil || len(data) == 0 {
			err = errors.New("no data")
			return
		}
		col = []string{}
		for k := range data[0] {
			col = append(col, k)
		}
	}
	title := sht.AddRow()
	for _, c := range col {
		if maps != nil {
			title.AddCell().Value = maps[c]
		} else {
			title.AddCell().Value = c
		}
	}
	// 添加一个空行
	sht.AddRow()

	for _, row := range data {
		newRow := sht.AddRow()
		for _, value := range col {
			if v, ok := row[value]; ok {
				newRow.AddCell().Value = fmt.Sprintf("%v", v)
			} else {
				newRow.AddCell().Value = ""
			}
		}
	}

	assetsAbsPath := beego.AppPath + "/" +
		beego.AppConfig.String("assets_save_path")
	keyword := c.GetString("query") + "|" + c.GetString("fields") + "|" + c.GetString("offset")
	fileAbsPath := assetsAbsPath + ParseHashStr(encoder.Md5String(keyword)) + ".xlsx"
	fDir := filepath.Dir(fileAbsPath)
	if _, e := os.Stat(fDir); os.IsNotExist(e) {
		e = os.MkdirAll(fDir, 0755)
		if e != nil {
			err = e
			return
		}
	}

	err = file.Save(fileAbsPath)
	if err != nil {
		return
	}
	defer func() {
		// delete file
		err := os.Remove(fileAbsPath)
		if err != nil {
			glog.Error("EXCEL", err)
		}
	}()
	c.Ctx.ResponseWriter.ResponseWriter.Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=%s", name+".xlsx"))
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, fileAbsPath)
	return
}
