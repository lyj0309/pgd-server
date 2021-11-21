package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pgd/controller"
	"pgd/global"
)


func Register(c *gin.Context) {
	form := make(map[string]interface{})
	err := c.ShouldBindJSON(&form)
	if err != nil {
		fmt.Println(err)
		controller.ReturnFalse(c)
		return
	}

	global.DB.Create(&global.User{
		ID:       form["id"].(string),
		Name:     form["name"].(string),
		Password: form["pwd"].(string),
	})
}
