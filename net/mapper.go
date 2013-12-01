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
	mapper map[string]Module
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

	//handle request
	err := module.Handle(request, response, conn)
	if err != nil {
		return err
	}
	response.Set("module", name)
	return nil
}
