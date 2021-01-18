/**
 * Author: Goddy <goddy@mykg.ai> 2021-01-15
 */
package entity

import "mykg.ai/rua/domain/embed"

type FolderTag struct {
	embed.BasicEmbed `gorm:"embedded"`

	FolderId uint64 `gorm:"uniqueIndex:idx_ft,priority:1;not null" json:"folder_id"`
	TagId    uint64 `gorm:"uniqueIndex:idx_ft,priority:2;not null" json:"tag_id"`
}

func (FolderTag) TableName() string {
	return "folder_tag"
}
