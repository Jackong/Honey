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
type WrapFunc func(Request, Response, *Conn)

func (this ModuleFunc) Handle(req Request, res Response, conn *Conn) error {
	return this(req, res, conn)
}

type WrapModule struct {
	beforeFuncs []WrapFunc
	afterFuncs []WrapFunc
	Module
}

func (this *WrapModule) Handle(request Request, response Response, conn *Conn) error {
	for _, beforeFunc := range this.beforeFuncs {
		beforeFunc(request, response, conn)
		if HasResult(response) {
			return nil
		}
	}

	err := this.Module.Handle(request, response, conn)
	if err != nil {
		return err
	}

	for _, afterFunc := range this.afterFuncs {
		afterFunc(request, response, conn)
	}
	return nil
}

func (this *WrapModule) Before(beforeHandlers ...WrapFunc) *WrapModule{
	this.beforeFuncs = append(this.beforeFuncs, beforeHandlers...)
	return this
}

func (this *WrapModule) After(afterHandlers ...WrapFunc) *WrapModule {
	this.afterFuncs = append(this.afterFuncs, afterHandlers...)
	return this
}
