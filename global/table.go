package global

import "time"
import "gorm.io/datatypes"

type Order struct {
	ID             int
	CreateTime     *time.Time
	Creator        string         //下单者
	CreatorName    string         //下单者姓名
	Data           datatypes.JSON //各种信息JSON
	Accepter       string         //接单者
	AccepterName   string         //接单者姓名
	CompleteTime   *time.Time
	CompleteDetail datatypes.JSON
}

type User struct {
	ID       string //学号
	Name     string
	Password string
}

type Notice struct {
	ID         int
	Title      string
	Data       string
	CreateTime *time.Time
}
