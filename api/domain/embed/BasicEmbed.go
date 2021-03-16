/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-09
 */
package embed

import (
	"fmt"
	"time"
)

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type BasicEmbed struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Creator   uint64    `json:"creator"`
}
