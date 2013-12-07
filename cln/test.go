/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午11:34
 */
package main

import (
	"github.com/Jackong/Honey/net"
	. "github.com/Jackong/Honey/global"
	"fmt"
	"github.com/Jackong/Honey/net/handler/json"
)

func main() {
	cln, err := net.NewClient("localhost" + Project.String("server", "addr"), &json.Handler{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write")
	request := json.NewRequest()
	request.Set("module", "signIn")
	request.Set("email", "jk7@gmail.com")
	request.Set("password", "01234567890123456789012345678901")
	request.Set("name", "jk6")
	err = cln.HandleWrite(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("read")
	response := json.NewResponse()
	cln.HandleRead(response)
	fmt.Println(response.Get("code"), response.Get("msg"))
}

