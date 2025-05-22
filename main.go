package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// db, err := gorm.Open(sqlite.Open("./main.db"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }

	// redis := redis.NewClient(&redis.Options{
	// 	Addr:         "localhost:6379", // macOS 本地 Redis
	// 	Password:     "",               // 無密碼
	// 	DB:           0,                // 預設 DB
	// 	PoolSize:     10,               // 連線池大小
	// 	MinIdleConns: 5,                // 最小閒置連線
	// })

	// _ = redis

	// r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "POST", "OPTIONS"},
	// 	AllowHeaders:     []string{"Content-Type", "User-Agent", "Accept", "Authorization"},
	// 	AllowCredentials: false,
	// 	MaxAge:           12 * 60 * 60,
	// }))

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "pong"})
	// })

	// r.POST("regist", func(c *gin.Context) { handler.Regist(c, db) })
	// r.POST("login", func(c *gin.Context) { handler.Login(c, db) })

	// protected := r.Group("/", middleware.JWTAuthMiddleware())
	// protected.GET("product", func(c *gin.Context) { handler.Index(c, db, redis) })
	// protected.GET("item", func(c *gin.Context) { handler.Item(c, db, redis) })
	// protected.GET("my", func(c *gin.Context) { handler.My(c, db) })

	// protected.POST("sale", func(c *gin.Context) { handler.Sale(c, db) })
	// protected.POST("remove", func(c *gin.Context) { handler.Remove(c, db) })

	// r.Run()

	inputFile := "item.json"
	outputFile := "data.json"
	includeKeywords := []string{"10%", "60%", "100%"}
	excludeKeywords := []string{"企鵝國王"}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("讀取檔案錯誤:", err)
		return
	}

	var inputMap map[string]string
	if err := json.Unmarshal(data, &inputMap); err != nil {
		fmt.Println("解析 JSON 錯誤:", err)
		return
	}

	filteredMap := make(map[string]string)
	for k, v := range inputMap {
		// 檢查是否包含 include 關鍵字
		include := false
		for _, keyword := range includeKeywords {
			if strings.Contains(v, keyword) {
				include = true
				break
			}
		}
		if !include {
			continue // 沒有包含關鍵字就跳過
		}

		// 檢查是否有排除字
		exclude := false
		for _, keyword := range excludeKeywords {
			if strings.Contains(v, keyword) {
				exclude = true
				break
			}
		}
		if exclude {
			continue // 有排除關鍵字就跳過
		}

		processed := strings.ReplaceAll(v, "、", "")
		processed = strings.ReplaceAll(processed, "%", "p")

		// 通過篩選，加進結果
		filteredMap[k] = v
	}

	outputData, err := json.MarshalIndent(filteredMap, "", "  ")
	if err != nil {
		fmt.Println("編碼 JSON 錯誤:", err)
		return
	}

	if err := os.WriteFile(outputFile, outputData, 0644); err != nil {
		fmt.Println("寫入檔案錯誤:", err)
		return
	}

	fmt.Println("篩選結果已寫入", outputFile)
}
