package models

import (
	"larago/pkg/types"
	"time"
)

// BaseModel 模型基类
type BaseModelId struct {
	//主键id
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`
}

type BaseModelTime struct {
	//创建时间
	CreatedAt time.Time `gorm:"column:created_at;index"`
	//更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModelId) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
