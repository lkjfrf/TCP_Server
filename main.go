package main

import (
	"encoding/json"
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

	jsonSize, err := conn.Read(recvBuffer)
	if jsonSize > 0 && err == nil {
		PacketLength := recvBuffer[0]
		PacketType := recvBuffer[2]
		//	data := string(recvBuffer[:jsonSize])
		data := recvBuffer[:PacketLength]
		fmt.Println(data)
		fmt.Println(PacketType)
		fmt.Println(PacketLength)
		data = data[4:]

		var dat map[string]interface{}
		if unmarshalErr := json.Unmarshal(data, &dat); unmarshalErr != nil {
			fmt.Println("Unmarshal fail")
			return
		}
		content.ContentManagerInst().CallBack(int(PacketType), &conn, dat)

		//content.ContentManagerInst().CallBack(int(PacketType))
		//conn.Write(recvBuffer[:jsonSize])
		//SendPacket(&conn)
	}

	// bufferedPeek, _ := recvBuffer.Peek(1)
	// 	fmt.Println(bufferedPeek)
}

func Connect() {
	Addr, err := net.Listen("tcp", ":1998")
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
