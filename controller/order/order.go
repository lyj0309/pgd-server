package order

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"pgd/controller"
	"pgd/global"
	"time"
)

func Handler(c *gin.Context) {

}

func Detail(c *gin.Context) {
	id := c.Query(`id`)
	if id == "" {
		controller.ReturnFalse(c)
		return
	}
	var order global.Order
	global.DB.Where(`id = ?`, id).Find(&order)
	controller.ReturnTrue(c, order)
}
func List(c *gin.Context) {
	var orders []global.Order
	switch c.Request.URL.Path[7:] {
	case `list`:
		global.DB.Where(`creator = ?`, c.GetString(`uid`)).Order("create_time DESC").Find(&orders)
	case `receive_list`:
		global.DB.Where(`accepter = ?`, c.GetString(`uid`)).Order("create_time DESC").Find(&orders)
	case `can_receive`:
		global.DB.Where(`accepter = ''`).Order("create_time DESC").Find(&orders)
	}
	controller.ReturnTrue(c, orders)
}
func New(c *gin.Context) {
	form := make(map[string]interface{})
	err := c.ShouldBindJSON(&form)
	if err != nil {
		controller.ReturnFalse(c)
		return
	}
	data, err := json.Marshal(form)
	if err != nil {
		controller.ReturnFalse(c)
		return
	}

	now := time.Now()
	global.DB.Create(&global.Order{
		CreateTime:  &now,
		Creator:     c.GetString(`uid`),
		CreatorName: c.GetString(`name`),
		Data:        data,
	})
	controller.ReturnTrue(c, "成功")
}

func Receive(c *gin.Context) {
	id := c.Query(`id`)
	if id == "" {
		controller.ReturnFalse(c)
	}
	global.DB.Model(&global.Order{}).Where(`id = ?`, id).Updates(
		global.Order{
			Accepter:     c.GetString(`uid`),
			AccepterName: c.GetString(`name`),
		},
	)
	controller.ReturnTrue(c, "成功")
}
func Delete(c *gin.Context) {
	id := c.Query(`id`)
	if id == "" {
		controller.ReturnFalse(c)
	}
	global.DB.Delete(&global.Order{}, id)
	controller.ReturnTrue(c, "成功")
}

func Complete(c *gin.Context) {
	id := c.Query(`id`)
	if id == "" {
		controller.ReturnFalse(c)
	}

	form := make(map[string]interface{})
	err := c.ShouldBindJSON(&form)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := json.Marshal(form)
	now := time.Now()
	global.DB.Model(&global.Order{}).Where(`id = ?`, id).Updates(
		global.Order{
			CompleteTime:   &now,
			CompleteDetail: data,
		})
	controller.ReturnTrue(c, "成功")
}

func Cancel(c *gin.Context) {
	id := c.Query(`id`)
	if id == "" {
		controller.ReturnFalse(c)
	}

	global.DB.Model(&global.Order{}).Where(`id = ?`, id).Updates(map[string]interface{}{
		"accepter":        "",
		"accepter_name":   "",
		"complete_time":   nil,
		"complete_Detail": nil,
	})
	controller.ReturnTrue(c, "成功")
}
