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

func (ct *CTManager) CallBack(PacketType int, c *net.Conn, v interface{}) {
	switch PacketType {
	case 1:
		type SignIn struct {
			Id         string
			IntData    int32
			StringData string
		}
		data := &SignIn{}
		helper.FillStruct_Interface(v, &data)
		data.IntData = 1111
		data.StringData = "packet1"

		ct.SendPacket(c, PacketType, data)

	case 2:
		data := &Packet{}
		helper.FillStruct_Interface(v, &data)
		data.Data = "songsongsong"
		data.Data2 = 123123

		ct.SendPacket(c, PacketType, data)

	default:
		type SignIn struct {
			Id         string
			IntData    int32
			StringData string
		}
		data := &SignIn{}
		helper.FillStruct_Interface(v, &data)
		data.IntData = 321321
		data.StringData = "songsongE"

		ct.SendPacket(c, PacketType, data)
	}
}

func (ct *CTManager) SendPacket(c *net.Conn, PacketType int, v interface{}) {
	go func() {
		if c != nil {

			str, err := json.Marshal(v)
			if err != nil {
				fmt.Println("Marshal faile")
			}

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
