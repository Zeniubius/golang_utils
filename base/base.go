package controllers

import (
	"github.com/astaxie/beego"
	"kuaifa.com/kuaifa/work-together/models/bean"
	"github.com/astaxie/beego/validation"
	"errors"
	"kuaifa.com/kuaifa/work-together/utils"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) RespJSON(code int, data interface{}) {
	c.AllowCross()
	c.Ctx.Output.SetStatus(code)
	var hasIndent = true
	if beego.BConfig.RunMode == beego.PROD {
		hasIndent = false
	}
	c.Ctx.Output.JSON(data, hasIndent, false)
}

// 只有数据, 返回值默认为200(成功)
func (c *BaseController) RespJSONData(data interface{}) {
	c.AllowCross()
	c.RespJSON(bean.CODE_Success, data)
}

// 只有数据, 返回值默认为200(成功)
func (c *BaseController) RespJSONDataWithTotal(data interface{}, total int64) {
	c.RespJSON(bean.CODE_Success, map[string]interface{}{
		"data":  data,
		"total": total,
	})
}

func (c *BaseController) RespJSONDataWithSumAndTotal(data interface{}, total int64, sum float64) {
	c.RespJSON(bean.CODE_Success, map[string]interface{}{
		"data":  data,
		"total": total,
		"sum": sum,
	})
}

func (c *BaseController) RespExcel(data []map[string]interface{}, name string, cols []string, maps map[string]string) {
	err := utils.ServerExcelMap(data, name, cols, maps, c.Controller)
	if err != nil {
		c.RespJSON(bean.CODE_Bad_Request, err.Error())
		return
	}
}

func (c *BaseController) Uid() int {
	return c.Ctx.Input.GetData("uid").(int)
}

func (c *BaseController) Validate(u interface{}) error {
	valid := validation.Validation{}
	b, err := valid.Valid(u)
	if err != nil {
		return err
	}
	if !b {
		e := valid.Errors[0]
		err = errors.New(e.Key + " " + e.Message)
		return err
	}
	return nil
}

func (c *BaseController) AllowCross() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")                               //允许访问源
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST,DELETE, GET, PUT, OPTIONS") //允许post访问
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")     //header的类型
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}

func (c *BaseController) Options() {
	c.AllowCross() //允许跨域
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}
