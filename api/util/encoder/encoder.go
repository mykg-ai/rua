/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-11
 */
package encoder

type Encoder interface {
	encrypt(str string) string
	decrypt(str string) string
}
