/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-09
 */
package embed

import "time"

type BasicEmbed struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Creator   string    `json:"creator"`
}
