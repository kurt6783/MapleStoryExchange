package main

import (
	"MapleStoryExchange/internal/handler"
	"MapleStoryExchange/internal/middleware"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

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

	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

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

func scraper() {
	minPage := 1
	maxPage := 37
	e := "it"

	var results []string
	fmt.Println("開始")
	for page := minPage; page <= maxPage; page++ {
		fmt.Println("第" + strconv.Itoa(page) + "頁")
		url := fmt.Sprintf("https://maple.yampiz.com/tw/p/strategy-item/it_?p=%d", page)
		res, err := http.Get(url)
		if err != nil {
			log.Printf("無法取得網頁：%v", err)
			continue
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Printf("HTTP 錯誤：%d", res.StatusCode)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Printf("解析 HTML 發生錯誤：%v", err)
			continue
		}

		doc.Find(".important").Each(func(i int, s *goquery.Selection) {
			name := strings.TrimSpace(s.Text())
			if name != "" {
				sql := fmt.Sprintf("INSERT INTO product (name, category) VALUES ('%s', '');", name)
				results = append(results, sql)
			}
		})

		time.Sleep(1 * time.Second) // 延遲一秒
	}
	outputPath := "/Users/kurt.hsu/Desktop/MapleStoryExchange/sql/" + e + ".sql"
	content := strings.Join(results, "\n") + "\n"
	err := os.WriteFile(outputPath, []byte(content), 0644)
	if err != nil {

	}
	fmt.Println("結束")
}
