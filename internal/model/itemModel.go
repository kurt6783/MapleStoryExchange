package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"MapleStoryExchange/internal/model/types"

	"gorm.io/gorm"
)

func NewItemModel(db *gorm.DB) *ItemModel {
	return &ItemModel{
		db,
	}
}

type ItemModel struct {
	db *gorm.DB
}

func (m *ItemModel) Item(ctx context.Context, productId int) ([]types.ItemWithName, error) {
	var items []types.ItemWithName
	err := m.db.WithContext(ctx).Raw(`
		SELECT i.id, i.product_id, p.name as name, i.price, i.memo, i.owner_id, u.code as owner_code
		FROM item i
		LEFT JOIN product p ON i.product_id = p.id
		LEFT JOIN user u ON i.owner_id = u.id
		WHERE i.product_id = ?
		ORDER BY i.price ASC
    `, productId).Scan(&items).Error
	if err != nil {
		return items, err
	}
	return items, nil
}

func (m *ItemModel) Sale(ctx context.Context, productId int, owner_id int, price int, memo string) (types.Item, error) {
	item := types.Item{
		ProductID: productId,
		OwnerID:   owner_id,
		Status:    true,
		Price:     price,
		Memo:      memo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := m.db.WithContext(ctx).Create(&item).Error; err != nil {
		return types.Item{}, err
	}
	return item, nil
}

func (m *ItemModel) My(ctx context.Context, userId int) ([]types.ItemWithName, error) {
	var items []types.ItemWithName
	err := m.db.WithContext(ctx).Raw(`
		SELECT i.id, i.product_id, p.name as name, i.price, i.memo, i.owner_id, u.code as owner_code
		FROM item i
		LEFT JOIN product p ON i.product_id = p.id
		LEFT JOIN user u ON i.owner_id = u.id
		WHERE u.id = ?
		ORDER BY i.price DESC
    `, userId).Scan(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (m *ItemModel) Remove(ctx context.Context, itemId int, owner_id int) (types.Item, error) {
	var item types.Item

	// 查找符合 productId 和 owner_id 的 item
	result := m.db.WithContext(ctx).
		Where("id = ? AND owner_id = ?", itemId, owner_id).
		First(&item)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return types.Item{}, fmt.Errorf("item with id %d and owner_id %d not found", itemId, owner_id)
		}
		return types.Item{}, result.Error
	}

	// 刪除找到的 item
	if err := m.db.WithContext(ctx).
		Where("id = ? AND owner_id = ?", itemId, owner_id).
		Delete(&types.Item{}).Error; err != nil {
		return types.Item{}, err
	}

	return item, nil
}
