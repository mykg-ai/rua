/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-09
 */
package entity

import "mykg.ai/rua/domain/embed"

type LinkTag struct {
	embed.BasicEmbed `gorm:"embedded"`

	LinkId uint64 `gorm:"uniqueIndex:idx_lt,priority:1;not null" json:"link_id"`
	TagId  uint64 `gorm:"uniqueIndex:idx_lt,priority:2;not null" json:"tag_id"`
}

func (LinkTag) TableName() string {
	return "link_tag"
}
