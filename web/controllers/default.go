package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
func (c *MainController) SetHouseInfo() {
	c.TplName = "setHouseInfo.html"
	fmt.Println("SetHouseInfo")

	// 获取用户输入
	rentingID := c.GetString("rentingID")
	fczbh := c.GetString("fczbh")
	fzxm := c.GetString("fzxm")
	djrq := c.GetString("djrq")
	zfmj := c.GetString("zfmj")
	fwsjyt := c.GetString("fwsjyt")
	sfdy := c.GetString("sfdy")

	if rentingID == "" {
		fmt.Println("rentingID不应为空!")
		return
	}

	// 组织参数
	var args []string
	args = append(args, "addHouseInfo")
	args = append(args, rentingID)
	args = append(args, fczbh)
	args = append(args, fzxm)
	args = append(args, djrq)
	args = append(args, zfmj)
	args = append(args, fwsjyt)
	args = append(args, sfdy)

	fmt.Printf("添加房屋信息输入的数据args :%s\n", args)

	//TODO 添加到fabric账本中

}
func (c *MainController) GetHouseInfo() {
	c.TplName = "getHouseInfo.html"
}
