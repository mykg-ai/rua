/**
 * Author: Goddy <goddy@mykg.ai> 2021-01-15
 */
package entity

import "mykg.ai/rua/domain/embed"

type FolderView struct {
	embed.BasicEmbed `gorm:"embedded"`

	Nid  uint64 `gorm:"uniqueIndex:idx_nn,priority:1;not null" json:"nid"`
	Name string `gorm:"uniqueIndex:idx_nn,priority:2;size:20;not null" json:"name"`
}

func (FolderView) TableName() string {
	return "folder_view"
}
