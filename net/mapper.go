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
	beforeFuncs []ModuleFunc
	mapper map[string]Module
	afterFuncs []ModuleFunc
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

func AllBefore(before...ModuleFunc) {
	beforeFuncs = append(beforeFuncs, before...)
}

func AllAfter(after...ModuleFunc) {
	afterFuncs = append(afterFuncs, after...)
}

func Handle(request Request, response Response, conn *Conn) (error){
	name := request.Get("module")
	defer func() {
		response.Set("module", name)
		if e := recover(); e != nil {
			rtErr := e.(err.Runtime)
			Log.Error(rtErr)
			response.Set("code", rtErr.Code)
			response.Set("msg", rtErr.Msg)
		}
	}()
	//check and get module
	if name == nil {
		panic(err.Runtime{Code: err.CODE_MODULE_NOT_FOUND, Msg: "module is required, but not found"})
	}
	module, ok := mapper[name.(string)]
	if !ok {
		panic(err.Runtime{Code: err.CODE_MODULE_NOT_FOUND, Msg: fmt.Sprintf("module %v not found", name)})
	}

	//before handler
	for _, beforeFunc := range beforeFuncs {
		err := beforeFunc(request, response, conn)
		if err != nil {
			return err
		}
	}

	//handle request
	err := module.Handle(request, response, conn)
	if err != nil {
		return err
	}

	//after handler
	for _, afterFunc := range afterFuncs {
		afterFunc(request, response, conn)
	}
	return nil
}
