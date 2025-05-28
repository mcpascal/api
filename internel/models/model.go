package models

type BaseModel struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:ID"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt int64  `json:"deleted_at,omitempty" gorm:"comment:删除时间"` // 软删除字段ßß
}

