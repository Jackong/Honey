/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 下午3:27
 */
package net

import (
	"github.com/Jackong/Honey/err"
	"fmt"
	"regexp"
)

type Protocol interface {
	Get(string) interface {}
	Set(string, interface {})
	Decode([]byte) (error)
	Encode() ([]byte, error)
}

type Request interface {
	Protocol
}

type Response interface {
	Protocol
}

func Required(protocol Protocol, key string) interface {}{
	val := protocol.Get(key)
	if val != nil {
		return val
	}
	panic(err.Runtime{Code: err.CODE_INPUT, Msg: fmt.Sprintf("The %v is required for %v", key, protocol.Get("module"))})
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
		msg = fmt.Sprintf("The %v must be pattern for %v", key, protocol.Get("module"))
	default:
		msg = fmt.Sprintf("The %v must be string type for %v", key, protocol.Get("module"))
	}
	panic(err.Runtime{Code: err.CODE_INPUT, Msg: msg})
}
