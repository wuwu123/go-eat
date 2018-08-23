package common

import (
	"github.com/go-redis/redis"
	"github.com/astaxie/beego"
	"fmt"
)

func GetRedis() *redis.Client {
	host := beego.AppConfig.String("cache_host")
	port := beego.AppConfig.DefaultInt("cache_port", 6379)
	password := beego.AppConfig.DefaultString("cache_pwd", "")
	database := beego.AppConfig.DefaultInt("cache_database", 0)
	if host == "" {
		beego.Error("redis host 无效->", host)
	}
	addr := fmt.Sprintf("%s:%d" , host , port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       database,
	})
	_, err := client.Ping().Result()
	if err != nil {
		beego.Error("redis 建立链接失败")
	}
	return client
}
