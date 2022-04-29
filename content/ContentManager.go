package content

import "sync"

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

}
