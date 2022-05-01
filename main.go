package main

import (
	"fmt"
	"main/content"
	"net"

	"github.com/mitchellh/mapstructure"
)

func FillStruct(data map[string]interface{}, result interface{}) {
	if err := mapstructure.Decode(data, &result); err != nil {
		fmt.Println(err)
	}
}

func Read(conn net.Conn) {
	recvBuffer := make([]byte, 1024) // 1024 == 1KB
	//SetWriteBuffer(2 * 1024 * 1024)

	for {

		go func() {
			jsonSize, err := conn.Read(recvBuffer)

			if jsonSize > 0 && err == nil {
				PacketLength := recvBuffer[0]
				PacketType := recvBuffer[3]

				//	data := string(recvBuffer[:jsonSize])
				data := recvBuffer[:PacketLength]

				fmt.Println(data)
				fmt.Println(PacketType)
				fmt.Println(PacketLength)

				//content.ContentManagerInst().CallBack(int(PacketType))

				//conn.Write(recvBuffer[:jsonSize])
				//SendPacket(&conn)
			}
		}()

		// bufferedPeek, _ := recvBuffer.Peek(1)
		// 	fmt.Println(bufferedPeek)
	}
}

func Connect() {
	Addr, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Connection Fail : ", err)
	}

	defer Addr.Close()

	for {
		conn, err := Addr.Accept()
		if err != nil {
			fmt.Println("Connect Fail : ", err)
		} else {
			fmt.Println("Connect Success : ", conn)
		}

		go Read(conn)
	}
}

func main() {
	fmt.Println("Server Start")
	content.ContentManagerInst()

	Connect()
}
