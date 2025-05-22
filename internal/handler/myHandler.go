package handler

import (
	"MapleStoryExchange/internal/model"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func My(c *gin.Context, db *gorm.DB) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "使用者資訊不存在"})
		return
	}

	uid, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "使用者 ID 型別錯誤"})
		return
	}

	itemModel := model.NewItemModel(db)
	items, err := itemModel.My(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch items: " + err.Error(),
		})
		return
	}

	// 序列化查詢結果為 JSON
	data, err := json.Marshal(gin.H{"data": items})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to marshal response: " + err.Error(),
		})
		return
	}

	// 返回結果
	c.Data(http.StatusOK, "application/json", data)
}
