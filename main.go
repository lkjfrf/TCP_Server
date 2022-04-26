package main

import (
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
			data := string(recvBuffer[:jsonSize])
			fmt.Println(data)
			//conn.Write(recvBuffer[:jsonSize])
			SendPacket(&conn)
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
			Length := uint8(len(Data) + 4)
			//sizeBytes := make([]byte, 4)
			var sizeBytes []byte
			//sizeBytes := []byte(uint8)

			PacketType := uint8(3)
			PacketHadder := []byte{byte(Length), 0, byte(PacketType), 0}

			//binary.LittleEndian.PutUint32(sizeBytes, Length)

			sizeBytes = append(sizeBytes, PacketHadder...)
			sizeBytes = append(sizeBytes, Data...)

			(*c).Write(sizeBytes)
		}
	}()
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

	Connect()
}
