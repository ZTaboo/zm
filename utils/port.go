package utils

import (
	"log"
	"net"
)

func PortCheck(port string) bool {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return true
	} else {
		listener, err := net.Listen("tcp4", ":"+port)
		if err != nil {
			return true
		} else {
			defer func(listen net.Listener) {
				err := listen.Close()
				if err != nil {
					log.Println("Listen close failed: ", err.Error())
					return
				}
			}(listen)
			defer func(listen net.Listener) {
				err := listen.Close()
				if err != nil {
					log.Println("Listen close failed: ", err.Error())
					return
				}
			}(listener)
			return false
		}

	}
}
