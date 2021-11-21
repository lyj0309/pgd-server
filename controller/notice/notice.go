package notice

import (
	"github.com/gin-gonic/gin"
	"pgd/controller"
	"pgd/global"
	"time"
)


func New(c *gin.Context) {
	form := make(map[string]interface{})
	err := c.ShouldBindJSON(&form)
	if err != nil {
		controller.ReturnFalse(c)
		return
	}
	now := time.Now()
	global.DB.Create(&global.Notice{
		Title:      form["title"].(string),
		Data:       form["data"].(string),
		CreateTime: &now,
	})
	controller.ReturnTrue(c,"成功")

}
func Delete(c *gin.Context) {
	global.DB.Where(`id = ?`,c.Query(`id`)).Delete(global.Notice{})
	controller.ReturnTrue(c,"成功")
}
func List(c *gin.Context) {
	var notices []global.Notice
	global.DB.Find(&notices)
	controller.ReturnTrue(c,notices)
}
