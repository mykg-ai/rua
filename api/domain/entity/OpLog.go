/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-11
 */
package entity

import (
	"mykg.ai/rua/domain/enum"
	"time"
)

type OpLog struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Creator   string    `json:"creator"`

	Nid     uint64      `gorm:"index:idx_no,priority:1" json:"nid"`
	OpType  enum.OpType `gorm:"index:idx_no,priority:2" json:"op_type"`
	Content string      `gorm:"size:255" json:"content"`
}

func (OpLog) TableName() string {
	return "log_op"
}
