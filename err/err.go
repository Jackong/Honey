/**
 * User: jackong
 * Date: 11/12/13
 * Time: 11:14 AM
 */
package err

type code int

const (
	CODE_OK = 0
	CODE_INPUT = (iota - 1)
	CODE_MODULE_NOT_FOUND
)

type Runtime struct {
	Code int
	Msg string
}

func (this Runtime) Error() string{
	return this.Msg
}

type System struct {
	Code int
	Msg string
}

func (this System) Error() string{
	return this.Msg
}
