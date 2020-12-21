/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-09
 */
package entity

import (
	"mykg.ai/rua/domain/embed"
)

type Link struct {
	embed.BasicEmbed `gorm:"embedded"`

	Nid         uint64 `gorm:"index:idx_nid;not null" json:"nid"`
	Short       string `gorm:"uniqueIndex:idx_s;size:20;not null" json:"short"`
	Target      string `gorm:"size:255;not null" json:"target"`
	Description string `gorm:"size:255;not null" json:"description"`
	Enable bool `gorm:"default:true" json:"enable"`
}

func (Link) TableName() string {
	return "link"
}
