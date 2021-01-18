/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-09
 */
package entity

import (
	"mykg.ai/rua/domain/enum"
	"time"
)

type User struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UID      string        `gorm:"uniqueIndex:id_uid;size:40;not null" json:"uid"`
	Username string        `gorm:"uniqueIndex:idx_u;size:20;not null" json:"username"`
	Password string        `gorm:"size:20;not null" json:"password"`
	Identity enum.Identity `gorm:"size:20;not null" json:"identity"`
}

func (User) TableName() string {
	return "user"
}
