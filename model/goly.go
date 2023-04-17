package model

type Goly struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	FullUrl   string `gorm:"unique" json:"fullUrl"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updatedAt"`
}
