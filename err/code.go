/**
 * User: Jackong
 * Date: 13-12-4
 * Time: 下午11:23
 */
package err

type code int

const (
	CODE_OK = 0
	CODE_INPUT = iota
	CODE_MODULE_NOT_FOUND
	CODE_USER_ACCOUNT_EXIST
)
