package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"varchar(200);not null" json:"name"`
	Email     string    `gorm:"varchar(200);not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Address   *Address  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"address"`
}
