package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

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
	recvBuffer := make([]byte, 1024) // 1024 == 1KB
	//SetWriteBuffer(2 * 1024 * 1024)

	for {
		jsonSize, err := conn.Read(recvBuffer)

		if jsonSize > 0 && err == nil {
			//packet := Packet{}

			data := string(recvBuffer[:jsonSize])
			fmt.Println(data)
			conn.Write(recvBuffer[:jsonSize])
			//SendPacket(&conn)
		}
	}
}

func SendPacket(c *net.Conn) {
	data := Packet{Id: "songsong", Data: "me", Data2: 123}
	str, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Marshal Err")
	}

	go func() {
		if c != nil {
			Data := []byte(str)
			Length := uint32(len(Data))
			sizeBytes := make([]byte, 4)
			PacketType := int32(1)
			PacketHadder := []byte{byte(PacketType), 0, byte(Length), 0}

			binary.LittleEndian.PutUint32(sizeBytes, Length)

			sizeBytes = append(sizeBytes, PacketHadder...)
			sizeBytes = append(sizeBytes, Data...)

			(*c).Write(sizeBytes)
		}
	}()
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
