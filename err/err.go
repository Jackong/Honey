/**
 * User: jackong
 * Date: 11/12/13
 * Time: 11:14 AM
 */
package err

type Runtime struct {
	Code code
	Msg string
}

func (this Runtime) Error() string{
	return this.Msg
}

type System struct {
	Code code
	Msg string
}

func (this System) Error() string{
	return this.Msg
}
