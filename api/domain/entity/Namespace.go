/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-08
 */
package entity

import (
	"mykg.ai/rua/domain/embed"
)

type Namespace struct {
	embed.BasicEmbed `gorm:"embedded"`

	Name string `gorm:"uniqueIndex:idx_n;size:20;not null" json:"name"`
}

func (Namespace) TableName() string {
	return "namespace"
}
