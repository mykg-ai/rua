/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-11
 */
package enum

type OpType string

const (
	/** namespace */
	APPEND_NS OpType = "APPEND_NS"

	/** user */
	APPEND_USER_NS      OpType = "APPEND_USER_NS"
	REMOVE_USER_NS      OpType = "REMOVE_USER_NS"
	UPDATE_USER_NS_ROLE OpType = "UPDATE_USER_NS_ROLE"

	/** link */
	APPEND_LINK        OpType = "APPEND_LINK"
	REMOVE_LINK        OpType = "REMOVE_LINK"
	UPDATE_LINK_TARGET OpType = "UPDATE_LINK_TARGET"
	UPDATE_LINK_TAG    OpType = "UPDATE_LINK_TAG"
	ENABLE_LINK        OpType = "ENABLE_LINK"
	UNABLE_LINK        OpType = "UNABLE_LINK"
)
