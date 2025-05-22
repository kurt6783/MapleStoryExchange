package types

import "time"

var _ WithTableName = (*User)(nil)

type User struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Account   string    `gorm:"column:account;type:varchar(255);not null;unique" json:"account"`
	Password  string    `gorm:"column:password;type:text;not null" json:"-"` // 不序列化到 JSON
	Code      string    `gorm:"column:code;type:text;not null" json:"code"`
	Role      string    `gorm:"column:role;type:varchar(50);default:user" json:"role"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
