/**
 * Author: Goddy <goddy@mykg.ai> 2021-01-15
 */
package entity

import "mykg.ai/rua/domain/embed"

type Folder struct {
	embed.BasicEmbed `gorm:"embedded"`

	FVid uint64 `gorm:"uniqueIndex:idx_fp,priority:1;not null" json:"f_vid"`
	Path string `gorm:"uniqueIndex:idx_fp,priority:2;size:255;not null" json:"path"`
}

func (Folder) TableName() string {
	return "folder"
}
