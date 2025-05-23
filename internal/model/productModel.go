package model

import (
	"context"

	"MapleStoryExchange/internal/model/types"

	"gorm.io/gorm"
)

func NewProductModel(db *gorm.DB) *ProductModel {
	return &ProductModel{
		db,
	}
}

type ProductModel struct {
	db *gorm.DB
}

func (m *ProductModel) Index(ctx context.Context) ([]types.ProductWithCount, error) {
	var products []types.ProductWithCount
	err := m.db.WithContext(ctx).Raw(`
        SELECT p.id, p.name, p.category, COALESCE(COUNT(i.id), 0) as count
        FROM product p
        LEFT JOIN item i ON p.id = i.product_id
        GROUP BY p.id, p.name, p.category
        ORDER BY count DESC
    `).Scan(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
