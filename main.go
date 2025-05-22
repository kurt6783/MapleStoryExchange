package main

import (
	"MapleStoryExchange/internal/handler"
	"MapleStoryExchange/internal/middleware"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./main.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	redis := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379", // macOS 本地 Redis
		Password:     "",               // 無密碼
		DB:           0,                // 預設 DB
		PoolSize:     10,               // 連線池大小
		MinIdleConns: 5,                // 最小閒置連線
	})

	_ = redis

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "User-Agent", "Accept", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * 60 * 60,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("regist", func(c *gin.Context) { handler.Regist(c, db) })
	r.POST("login", func(c *gin.Context) { handler.Login(c, db) })

	protected := r.Group("/", middleware.JWTAuthMiddleware())
	protected.GET("product", func(c *gin.Context) { handler.Index(c, db, redis) })
	protected.GET("item", func(c *gin.Context) { handler.Item(c, db, redis) })
	protected.GET("my", func(c *gin.Context) { handler.My(c, db) })

	protected.POST("sale", func(c *gin.Context) { handler.Sale(c, db) })
	protected.POST("remove", func(c *gin.Context) { handler.Remove(c, db) })

	r.Run()
}
