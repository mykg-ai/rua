/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-11
 */
package entity

import "time"

type UsageLog struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Creator   string    `json:"creator"`

	Lid       uint64 `gorm:"index:idx_l" json:"lid"`
	Ip        string `gorm:"size:20" json:"ip"`
	UserAgent string `gorm:"size:200" json:"user_agent"`
}

func (UsageLog) TableName() string {
	return "log_usage"
}
