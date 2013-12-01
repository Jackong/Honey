/**
 * User: Jackong
 * Date: 13-11-30
 * Time: 下午9:52
 */
package json

import (
	"github.com/Jackong/Honey/net"
	"encoding/json"
	"errors"
)

type Handler struct {

}

type header struct {
	Len uint `json:"len"`
}

func (this *Handler) HeaderLength() int {
	return 12
}

func (this *Handler) FormatProtocol(protocol net.Protocol) ([]byte, error) {
	resBuf, err := protocol.Encode()
	if err != nil {
		return nil, err
	}
	header := header{Len: uint(len(resBuf))}
	headerBuf, err := json.Marshal(header)
	if err != nil {
		return nil, err
	}
	for index := len(headerBuf); index < this.HeaderLength(); index++ {
		headerBuf = append(headerBuf, '*')
	}
	return append(headerBuf, resBuf...), nil
}

func (this *Handler) HandleHeader(buf []byte) (uint, error) {
	end := 0
	for index, b := range buf {
		if b == '}' {
			end = index + 1
			break
		}
	}
	header := header{}
	err := json.Unmarshal(buf[0: end], &header)
	if err != nil {
		return 0, err
	}
	return header.Len, nil
}

func (this *Handler) HandleRequest(buf []byte, conn *net.Conn) ([]byte, error) {
	request := NewRequest()
	if err := request.Decode(buf); err != nil {
		return nil, err
	}
	name := request.Get("module")
	if name == nil {
		return nil, errors.New("request module not set")
	}
	module := net.GetModule(name.(string))
	if module == nil {
		return nil, errors.New("request module not found " + name.(string))
	}
	response := NewResponse()
	module.Handle(request, response, conn)
	respBuf, err := this.FormatProtocol(response)
	if err != nil {
		return nil, err
	}
	return respBuf, nil
}

func (this *Handler) HandleAcceptError(error) {

}
func (this *Handler) HandleConnError(interface {}) {

}
