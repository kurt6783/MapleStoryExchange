package handler

import (
	"MapleStoryExchange/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Regist(c *gin.Context, db *gorm.DB) {
	var req RegisterReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON format: " + err.Error(),
		})
		return
	}

	userModel := model.NewUserModel(db)
	user, err := userModel.Register(c, req.Account, req.Password, req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"account": user.Account,
			"code":    user.Code,
		},
	})
}

type RegisterReq struct {
	Account         string `json:"account" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
	Code            string `json:"code" binding:"required"`
}
