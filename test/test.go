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
	request.Set("module", "feedback")
	request.Set("msg", "ahddd")
	err = cln.HandleWrite(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("read")
	response := json.NewResponse()
	cln.HandleRead(response)
}

