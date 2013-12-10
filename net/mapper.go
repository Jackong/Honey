/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午9:27
 */
package net

import (
	"errors"
	"github.com/Jackong/Honey/err"
	. "github.com/Jackong/Honey/global"
	"fmt"
)

var (
	beforeFuncs []WrapFunc
	mapper map[string]Module
	afterFuncs []WrapFunc
)

func init() {
	mapper = make(map[string]Module)
}

func Attach(name string, module Module) *WrapModule {
	if module == nil {
		panic(errors.New("nil module for " + name))
	}
	if _, ok := mapper[name]; ok {
		panic(errors.New("Duplicated module for " + name))
	}
	wrap := &WrapModule{Module: module}
	mapper[name] = wrap
	return wrap
}

func AttachFunc(name string, module ModuleFunc) *WrapModule {
	return Attach(name, module)
}

func AllBefore(before...WrapFunc) {
	beforeFuncs = append(beforeFuncs, before...)
}

func AllAfter(after...WrapFunc) {
	afterFuncs = append(afterFuncs, after...)
}

func Handle(request Request, response Response, conn *Conn) {
	name := Pattern(request, "module", "[a-zA-Z0-9]{2,}")
	defer func() {
		response.Set("module", name)
		if e := recover(); e != nil {
			handleError(e, name, conn, response)
		}
	}()
	//check module
	module, ok := mapper[name]
	if !ok {
		msg := fmt.Sprintf("Module %v not found", name)
		Log.Alert("%v|%v", conn.Id, msg)
		Result(response, CODE_MODULE_NOT_FOUND, msg)
		return
	}

	//before handler
	for _, beforeFunc := range beforeFuncs {
		beforeFunc(request, response, conn)
		if HasResult(response) {
			return
		}
	}

	//handle request
	err := module.Handle(request, response, conn)
	if err != nil {
		panic(err)
		return
	}

	//after handler
	for _, afterFunc := range afterFuncs {
		afterFunc(request, response, conn)
	}
}

func handleError(e interface {}, module interface {}, conn *Conn, response Response) {
	switch val := e.(type){
	case err.Runtime:
		Result(response, val.Code, val.Msg, fmt.Sprintf("module:%v|%v", module, conn.Id))
	default:
		Result(response, CODE_UNKNOWN_ERROR, "unknown error on server", val)
	}
}
