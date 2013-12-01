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

func GetModule(name string) Module {
	return mapper[name]
}
