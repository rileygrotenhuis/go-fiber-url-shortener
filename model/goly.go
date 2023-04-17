package model

type Goly struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	ShortUrl  string `gorm:"unique;not null" json:"shortUrl"`
	FullUrl   string `gorm:"unique;not null" json:"fullUrl"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updatedAt"`
}
