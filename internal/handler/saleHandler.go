package handler

import (
	"MapleStoryExchange/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Sale(c *gin.Context, db *gorm.DB) {
	var req SaleReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON format: " + err.Error(),
		})
		return
	}

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
	_, err := itemModel.Sale(c, req.ProductId, uid, req.Price, req.Memo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{},
	})
}

type SaleReq struct {
	ProductId int    `json:"productId" binding:"required"`
	Price     int    `json:"price" binding:"required"`
	Memo      string `json:"memo" binding:"required"`
}
