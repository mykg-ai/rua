/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-09
 */
package entity

import (
	"mykg.ai/rua/domain/embed"
	"mykg.ai/rua/domain/enum"
)

type UserNamespace struct {
	embed.BasicEmbed `gorm:"embedded"`

	Uid  uint64    `gorm:"uniqueIndex:idx_un,priority:1;not null" json:"uid"`
	Nid  uint64    `gorm:"uniqueIndex:idx_un,priority:2;not null" json:"nid"`
	Role enum.Role `gorm:"size:20;not null" json:"role"`
}

func (UserNamespace) TableName() string {
	return "user_namespace"
}
