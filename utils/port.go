package utils

import (
	"log"
	"net"
)

// PortCheck 验证端口是否占用
func PortCheck(port string) bool {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		//fmt.Println("port:", port)
		//fmt.Println("listen failed: ", err.Error())
		return true
	} else {
		defer func(listen net.Listener) {
			err := listen.Close()
			if err != nil {
				log.Println("Listen close failed: ", err.Error())
				return
			}
		}(listen)
		return false
	}

}
