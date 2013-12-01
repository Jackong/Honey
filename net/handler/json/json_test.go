/**
 * User: Jackong
 * Date: 13-11-30
 * Time: 下午10:03
 */
package json

import (
	"testing"
	"encoding/json"
)

func TestHeader(t *testing.T) {
	buf := make([]byte, 15)
	buf = []byte(`{"len":9999}`)
	for index, b := range buf {
		if b == '}' {
			t.Log(index)
		}
	}
	t.Log(string(buf), len(buf))
	m := header{}
	err := json.Unmarshal(buf, &m)
	if err != nil {
		t.Error(err)
	}
	t.Log(m)
}

func TestJsonHeader(t *testing.T) {
	buf := []byte(`{"len":122}*`)
	t.Log(len(buf))
	hnd := &header{}
	length, err := hnd.HandleHeader(buf)
	if err != nil {
		t.Error(err)
	}
	t.Log(length)
	if length != 122 {
		t.Error("length error")
	}
}
