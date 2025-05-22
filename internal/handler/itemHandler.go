package handler

import (
	"MapleStoryExchange/internal/model"
	"MapleStoryExchange/internal/model/types"
	"context"
	"time"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Item(c *gin.Context, db *gorm.DB, redis *redis.Client) {
	// 定義快取鍵
	productIdStr := c.Query("productId")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid productId",
		})
		return
	}
	var cacheKey = "item" + productIdStr
	ctx := context.Background()

	// 檢查 Redis 快取
	cached, err := redis.Get(ctx, cacheKey).Result()
	if err == nil {
		// 快取命中，直接返回
		c.Data(http.StatusOK, "application/json", []byte(cached))
		return
	}

	// 快取未命中，查詢資料庫
	itemModel := model.NewItemModel(db)
	items, err := itemModel.Item(c, productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch items: " + err.Error(),
		})
		return
	}

	if items == nil {
		items = []types.ItemWithName{}
	}

	// 序列化查詢結果為 JSON
	data, err := json.Marshal(gin.H{"data": items})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to marshal response: " + err.Error(),
		})
		return
	}

	// 存入 Redis，設置 TTL 為 1 分鐘
	err = redis.Set(ctx, cacheKey, data, time.Minute).Err()
	if err != nil {
		// 記錄 Redis 錯誤，但不影響響應
		c.Error(err)
	}

	// 返回結果
	c.Data(http.StatusOK, "application/json", data)
}
