/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午10:20
 */
package net

type Module interface {
	Handle(Request, Response, *Conn) error
}

type ModuleFunc func(Request, Response, *Conn) error

func (this ModuleFunc) Handle(req Request, res Response, conn *Conn) error {
	return this(req, res, conn)
}

type WrapModule struct {
	beforeFuncs []ModuleFunc
	afterFuncs []ModuleFunc
	Module
}

func (this *WrapModule) Handle(req Request, resp Response, conn *Conn) error {
	for _, beforeFunc := range this.beforeFuncs {
		err := beforeFunc(req, resp, conn)
		if err != nil {
			return err
		}
	}

	err := this.Module.Handle(req, resp, conn)
	if err != nil {
		return err
	}

	for _, afterFunc := range this.afterFuncs {
		afterFunc(req, resp, conn)
	}
	return nil
}

func (this *WrapModule) Before(beforeHandlers ...ModuleFunc) *WrapModule{
	this.beforeFuncs = append(this.beforeFuncs, beforeHandlers...)
	return this
}

func (this *WrapModule) After(afterHandlers ...ModuleFunc) *WrapModule {
	this.afterFuncs = append(this.afterFuncs, afterHandlers...)
	return this
}
