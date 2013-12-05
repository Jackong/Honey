/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 下午3:27
 */
package net

import (
	"github.com/Jackong/Honey/err"
	. "github.com/Jackong/Honey/global"
	"fmt"
	"regexp"
)

type Protocol interface {
	Get(string) interface {}
	Set(string, interface {}) Protocol
	Decode([]byte) (error)
	Encode() ([]byte, error)
}

type Request interface {
	Protocol
}

type Response interface {
	Protocol
}

func Result(protocol Protocol, code int, msg string, ext ... interface {}) {
	protocol.Set("code", code).Set("msg", msg)
	info := fmt.Sprintf("%v:%v|%v", code, msg, ext)
	switch code {
	case CODE_OK:
		Log.Info(info)
	case CODE_UNKNOWN_ERROR:
		Log.Alert(info)
	default:
		Log.Error(info)
	}
}

func HasResult(protocol Protocol) bool {
	return protocol.Get("code") != nil
}

func Required(protocol Protocol, key string) interface {}{
	val := protocol.Get(key)
	if val != nil {
		return val
	}
	panic(err.Runtime{Code: CODE_INPUT, Msg: fmt.Sprintf("The %v is required", key)})
}

func Default(protocol Protocol, key string, value interface {}) interface {}{
	val := protocol.Get(key)
	if val != nil {
		return val
	}
	return value
}

func Pattern(protocol Protocol, key, pattern string) string {
	val := Required(protocol, key)
	var msg string
	switch value := val.(type) {
	case string:
		if match, _ := regexp.MatchString(pattern, value); match {
			return value
		}
		msg = fmt.Sprintf("The %v must be pattern", key)
	default:
		msg = fmt.Sprintf("The %v must be string type", key)
	}
	panic(err.Runtime{Code: CODE_INPUT, Msg: msg})
}
