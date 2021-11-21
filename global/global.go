package global

import (
	"fmt"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

var (
	DB              *gorm.DB
	Settings        map[string]interface{}
	OfficialAccount *officialaccount.OfficialAccount
)

func init() {
	var err error

	var b []byte
	b, err = ioutil.ReadFile(`setting.yml`)
	if err != nil {
		b, err = ioutil.ReadFile(`../setting.yml`)
		if err != nil {
			log.Fatalf("配置文件读取错误: %v", err)
		}
	}

	// 存储解析数据
	Settings = make(map[string]interface{})
	// 执行解析
	err = yaml.Unmarshal(b, &Settings)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(Settings)

	DB, err = gorm.Open(mysql.Open(Settings["mysql_dsn_local"].(string)), &gorm.Config{PrepareStmt: true}) //Logger: logger.Discard,

	wc := wechat.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:        Settings["redis_addr"].(string),
		Database:    0,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60, //second
	}
	redisCache := cache.NewRedis(redisOpts)

	OfficialAccount = wc.GetOfficialAccount(&offConfig.Config{
		AppID:          Settings["off_app_id"].(string),
		AppSecret:      Settings["off_app_secret"].(string),
		Token:          Settings["off_token"].(string),
		EncodingAESKey: Settings["off_encoding_AESKey"].(string),
		Cache:          redisCache,
	})

	if err != nil {
		fmt.Println(`数据库初始化错误`, err)
	}
}
