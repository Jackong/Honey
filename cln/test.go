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
	cln, err := net.NewClient("localhost" + Project.String("server", "addr"), Handler)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write")
	request := json.NewRequest()
	request.Set("module", "signUp")
	request.Set("email", "jk6@gmail.com")
	request.Set("password", "6pswd")
	request.Set("name", "jk6")
	err = cln.HandleWrite(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("read")
	response := json.NewResponse()
	cln.HandleRead(response)
	fmt.Println(response.Get("code"), response.Get("tips"))
}

