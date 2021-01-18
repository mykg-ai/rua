/**
 * Author: Goddy <goddy@mykg.ai> 2021-01-15
 */
package Recode

var (
	SUCCESS = Recode{Code: 0, Msg: "ok"}

	NOT_FOUND = Recode{Code: 404, Msg: "not found"}
)

type Recode struct {
	Code uint
	Msg  string
}
