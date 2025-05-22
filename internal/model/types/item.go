package types

import "time"

var _ WithTableName = (*Item)(nil)

type Item struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ProductID int       `gorm:"column:product_id;not null;index:idx_item_product_id" json:"product_id"`
	OwnerID   int       `gorm:"column:owner_id;not null;index:idx_item_owner_id" json:"owner_id"`
	Status    bool      `gorm:"column:status;default:true" json:"status"`
	Price     int       `gorm:"column:price;not null;check:price >= 0" json:"price"`
	Memo      string    `gorm:"column:memo;type:text" json:"memo"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type ItemWithName struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ProductID int       `gorm:"column:product_id;not null;index:idx_item_product_id" json:"product_id"`
	Name      string    `gorm:"column:name" json:"name"`
	OwnerID   int       `gorm:"column:owner_id;not null;index:idx_item_owner_id" json:"owner_id"`
	OwnerCode string    `gorm:"column:owner_code" json:"owner_code"`
	Status    bool      `gorm:"column:status;default:true" json:"status"`
	Price     int       `gorm:"column:price;not null;check:price >= 0" json:"price"`
	Memo      string    `gorm:"column:memo;type:text" json:"memo"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Item) TableName() string {
	return "item"
}
