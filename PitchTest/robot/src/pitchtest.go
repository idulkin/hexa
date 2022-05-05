package PitchTest

import (
	"mind/core/framework/skill"
	"mind/core/framework/drivers/hexabody"
    "time"
)

type PitchTest struct {
	skill.Base
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &PitchTest{}
}

func (d *PitchTest) OnStart() {
	// Use this method to do something when this skill is starting.
    hexabody.Start()
    hexabody.Stand()
}

func (d *PitchTest) OnClose() {
	// Use this method to do something when this skill is closing.
    hexabody.Close()
}

func (d *PitchTest) OnConnect() {
	// Use this method to do something when the remote connected.
    time.Sleep(time.Second * 3)
    hexabody.RotateHeadContinuously(1, 160)

    angle := -15.0

    for {
        for angle < 15.0 {
            hexabody.Pitch(angle, 300)
            angle++
        }

        for angle > -15.0 {
            hexabody.Pitch(angle, 300)
            angle--
        }
    }
}

func (d *PitchTest) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
    hexabody.Relax()
}

func (d *PitchTest) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *PitchTest) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}
