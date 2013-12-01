/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午9:27
 */
package net

import (
	"errors"
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
	//check and get module
	name := request.Get("module")
	if name == nil {
		return errors.New("request module not set")
	}
	module, ok := mapper[name.(string)]
	if !ok {
		return errors.New("request module not found " + name.(string))
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

	response.Set("module", name)
	return nil
}
