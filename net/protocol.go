/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 下午3:27
 */
package net

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
