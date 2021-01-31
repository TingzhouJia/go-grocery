package model

import "time"

type User struct {
	UserId string `gorm:"column:user_id" json:"userId"`
	NickName string `gorm:"column:nick_name" json:"nickName"`
	Mobile string `json:"mobile" gorm:"column:mobile" binding:"required"`
	Password string `json:"-" gorm:"column: password"`
	Address string `json:"address" gorm:"column:address"`
	IsDeleted bool `gorm:"column:is_deleted" json:"is_deleted"`
	IsLocked bool `json:"isLocked" gorm:"column:is_locked"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" gorm:"column:updated_at"`
}


