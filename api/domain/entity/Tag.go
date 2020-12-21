/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-09
 */
package entity

import "mykg.ai/rua/domain/embed"

type Tag struct {
	embed.BasicEmbed `gorm:"embedded"`

	Name string	`gorm:"uniqueIndex:idx_name;size:20;not null" json:"name"`
	Nid uint64	`gorm:"index:idx_nid;not null" json:"nid"`
}

func (Tag) TableName() string {
	return "tag"
}