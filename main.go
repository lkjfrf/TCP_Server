package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Packet struct {
	Id    string
	Data  string
	Data2 int32
}

func FillStruct(data map[string]interface{}, result interface{}) {
	if err := mapstructure.Decode(data, &result); err != nil {
		fmt.Println(err)
	}
}

func Read(conn net.Conn) {
	recvBuffer := make([]byte, 4096) // 1024 X 4 인데 1024 == 1KB 가 모여서 1MB 가됨
	var dat map[string]interface{}

	for {
		json, err := conn.Read(recvBuffer)

		if json > 0 && err == nil {
			//packet := Packet{}

			data := recvBuffer[:json]
			fmt.Println(string(data))

			splitedStrs := strings.Split(string(data), "\n")

			packet := Packet{}
			FillStruct(dat, &packet)

			for _, v := range splitedStrs {
				if len(v) == 0 {
					continue
				}

			}

			fmt.Println(packet.Data)
		}
	}
}

func main() {
	fmt.Println("INIT_Main")

	//	addAddr, err := net.ResolveTCPAddr("tcp", ":8000")

	Addr, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Connection Faile : ", err)
	}

	defer Addr.Close()

	for {
		conn, err := Addr.Accept()
		if err != nil {
			fmt.Println("Accept Faile : ", err)
		} else {
			fmt.Println("Connect Success : ", conn)
		}

		go Read(conn)
	}

}
