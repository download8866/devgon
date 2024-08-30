package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v9"
)

// var ctx = context.Background()

func main() {
	r := gin.Default()

	// 初始化Redis客户端
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "redis:6379", // Redis服务的地址, 在Docker中服务名称就是主机名
	// 	Password: "",           // 密码，如果没有可以忽略
	// 	DB:       0,            // 默认数据库
	// })

	r.GET("/", func(c *gin.Context) {
		fmt.Println("Service start. from path home ")
		// err := rdb.Set(ctx, "key", "value", 0).Err()
		// if err != nil {
		// 	log.Fatalf("Failed to connect to Redis: %v", err)
		// }

		// val, err := rdb.Get(ctx, "key").Result()
		// if err != nil {
		// 	log.Fatalf("Failed to get value from Redis: %v", err)
		// }

		// c.JSON(http.StatusOK, gin.H{
		// 	"message":     "Hello, World!",
		// 	"redis_value": val,
		// })
		c.JSON(200, gin.H{
			"message": "the redis not support , hot reloading. update5 ",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Println("Service start.")
	r.Run(":8080") // 在8080端口启动Gin服务器
}
