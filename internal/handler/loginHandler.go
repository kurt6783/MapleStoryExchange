package handler

import (
	"net/http"
	"time"

	"MapleStoryExchange/internal/model"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtKey = []byte("your-secret-key2")

type Claims struct {
	Account string `json:"account"`
	UserId  int    `json:"userId"`
	Code    string `json:"code"`
	jwt.RegisteredClaims
}

func Login(c *gin.Context, db *gorm.DB) {
	var req struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "參數錯誤"})
		return
	}

	userModel := model.NewUserModel(db)
	user, err := userModel.Find(c, req.Account)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "帳號不存在"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密碼錯誤"})
		return
	}

	// 產生 JWT
	expirationTime := time.Now().Add(365 * 24 * time.Hour)
	claims := &Claims{
		Account: user.Account,
		UserId:  int(user.ID),
		Code:    user.Code,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token 簽發失敗"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"account": user.Account,
			"code":    user.Code,
			"role":    user.Role,
		},
	})
}
