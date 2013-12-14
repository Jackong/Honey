/**
 * User: Jackong
 * Date: 13-12-5
 * Time: 下午8:59
 */
package global

const (
	CODE_UNKNOWN_ERROR = -1
	CODE_OK = (iota - 1)
	CODE_FAILED
	CODE_INPUT
	CODE_MODULE_NOT_FOUND
)

const (
	TIPS_NIL = iota
	TIPS_LABEL
	TIPS_FLOAT
	TIPS_ALERT
	TIPS_CONFIRMATION
)
