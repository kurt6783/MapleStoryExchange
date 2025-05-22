package model

import (
	"context"
	"time"

	"MapleStoryExchange/internal/model/types"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db,
	}
}

type UserModel struct {
	db *gorm.DB
}

func (m *UserModel) Register(ctx context.Context, account string, password string, code string) (types.User, error) {
	bcryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := types.User{
		Account:   account,
		Password:  string(bcryptedPassword),
		Code:      code,
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := m.db.WithContext(ctx).Create(&user).Error; err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (m *UserModel) Find(ctx context.Context, account string) (types.User, error) {
	var user types.User

	if err := m.db.WithContext(ctx).Model(&types.User{}).
		Where("account = ?", account).
		First(&user).Error; err != nil {
		return types.User{}, err
	}
	return user, nil
}
