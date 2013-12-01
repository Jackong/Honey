/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午11:28
 */
package json

import (
	"encoding/json"
)

type Protocol struct {
	values map[string]interface {}
}

func NewProtocol() *Protocol{
	return &Protocol{values: make(map[string] interface {})}
}

func (this *Protocol) Get(name string) interface {}{
	return this.values[name]
}

func (this *Protocol) Set(name string, value interface {}) {
	this.values[name] = value
}


func (this *Protocol) Decode(buf []byte) (error) {
	if err := json.Unmarshal(buf, &this.values);err != nil {
		return err
	}
	return nil
}

func (this *Protocol) Encode() ([]byte, error) {
	buf, err := json.Marshal(this.values)
	if err != nil {
		return nil, err
	}
	return buf, nil
}


type Request struct {
	*Protocol
}

func NewRequest() *Request {
	return &Request{Protocol: NewProtocol()}
}

type Response struct {
	*Protocol
}

func NewResponse() *Response {
	return &Response{Protocol: NewProtocol()}
}
