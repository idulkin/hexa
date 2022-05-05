package helloworld

import (
	"mind/core/framework/skill"
)

type helloworld struct {
	skill.Base
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &helloworld{}
}

func (d *helloworld) OnStart() {
	// Use this method to do something when this skill is starting.
}

func (d *helloworld) OnClose() {
	// Use this method to do something when this skill is closing.
}

func (d *helloworld) OnConnect() {
	// Use this method to do something when the remote connected.
}

func (d *helloworld) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d *helloworld) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *helloworld) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}
