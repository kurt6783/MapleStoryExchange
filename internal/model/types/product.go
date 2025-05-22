package types

var _ WithTableName = (*Product)(nil)

type Product struct {
	ID       uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Category string `gorm:"column:category" json:"category"`
	Count    int    `gorm:"-" json:"count"` // 不對應資料庫欄位，用於儲存 item 計數
}

type ProductWithCount struct {
	ID       uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Category string `gorm:"column:category" json:"category"`
	Count    int    `gorm:"column:count" json:"count"`
}

func (Product) TableName() string {
	return "product"
}
