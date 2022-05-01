package content

import (
	"encoding/json"
	"fmt"
	"main/helper"
	"net"
	"sync"
)

type Packet struct {
	Id    string
	Data  string
	Data2 int32
}

type CTManager struct {
	Callbacks map[string]func(interface{})
}

var instance_CT *CTManager
var once_CT sync.Once

func ContentManagerInst() *CTManager {
	once_CT.Do(func() {
		instance_CT = &CTManager{}
	})
	return instance_CT
}

func (ct *CTManager) Init() {
	fmt.Println("INIT_ContentManager")

}

func (ct *CTManager) CallBack(PacketType int, v interface{}) {
	switch PacketType {
	case 1:
		type SignIn struct {
			Conn *net.TCPConn
			Id   string
			Data int32
		}
		data := &SignIn{}
		helper.FillStruct_Interface(v, data)
		data.Data = 123123

		str, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Marshal faile")
		}

		ct.SendPacket(data.Conn, str, PacketType)

	case 2:
		type SignIn struct {
			Conn  *net.TCPConn
			Id    string
			Data2 int32
		}
		data := &SignIn{}
		helper.FillStruct_Interface(v, data)
		data.Data2 = 321321

		str, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Marshal faile")
		}

		ct.SendPacket(data.Conn, str, PacketType)
	}

}

func (ct *CTManager) SendPacket(c *net.TCPConn, str []byte, PacketType int) {
	go func() {
		if c != nil {
			Data := []byte(str)
			Length := uint8(len(Data) + 4)
			var sizeBytes []byte

			PacketType := uint8(PacketType)
			PacketHadder := []byte{byte(Length), 0, byte(PacketType), 0}

			//binary.LittleEndian.PutUint32(sizeBytes, Length)

			sizeBytes = append(sizeBytes, PacketHadder...)
			sizeBytes = append(sizeBytes, Data...)

			(*c).Write(sizeBytes)
		}
	}()
}
