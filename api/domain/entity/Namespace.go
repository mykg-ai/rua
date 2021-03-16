/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-08
 */
package entity

import (
	"encoding/json"
	"mykg.ai/rua/domain/embed"
)

type Namespace struct {
	embed.BasicEmbed `gorm:"embedded"`

	Name string `gorm:"uniqueIndex:idx_n;size:20;not null" json:"name"`
}

func (Namespace) TableName() string {
	return "namespace"
}

func (t Namespace) MarshalJSON() ([]byte, error) {
	type TmpJSON Namespace
	return json.Marshal(&struct {
		TmpJSON
		CreatedAt embed.DateTime `json:"created_at"`
		UpdatedAt embed.DateTime `json:"updated_at"`
	}{
		TmpJSON:   (TmpJSON)(t),
		CreatedAt: embed.DateTime(t.CreatedAt),
		UpdatedAt: embed.DateTime(t.UpdatedAt),
	})
}
