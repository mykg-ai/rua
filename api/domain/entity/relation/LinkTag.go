/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-09
 */
package entity

import "mykg.ai/rua/domain/embed"

type LinkTag struct {
	embed.BasicEmbed `gorm:"embedded"`

	Lid uint64	`gorm:"index:idx_lt,priority:1;not null" json:"lid"`
	Tid uint64	`gorm:"index:idx_lt,priority:2;not null" json:"tid"`
}

func (LinkTag) TableName() string {
	return "link_tag"
}