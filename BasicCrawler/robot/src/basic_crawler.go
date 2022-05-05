package SensorWalkSkill

import (
    "math"
    "mind/core/framework/drivers/distance"
    "mind/core/framework/drivers/hexabody"
    "mind/core/framework/log"
	"mind/core/framework/skill"
    "time"
)

const (
    WALK_SPEED          = 1.0 // cm per seccond
)

type SensorWalkSkill struct {
	skill.Base
}

func newDirection(direction float64, degrees float64) float64 {
    return math.Mod(direction+degrees, 360) * -1
}

/* Walk continuously, try not to hit anything */
func MoveAndScan() {
    direction := 0.0
    moveCount := 0
    time.Sleep(time.Second)
    hexabody.MoveHead(0, 0)
    hexabody.WalkContinuously(0, 0.5)
    hexabody.MoveHead(0, 0)
    hexabody.WalkContinuously(0, WALK_SPEED)

    for {
        dist, _ := distance.Value()

        // Something's too close, rotate and find another path
        if dist < 500 {
            hexabody.StopWalkingContinuously()
            angle := dist * 2
            direction = newDirection(direction, angle)
            hexabody.MoveHead(direction, 0)
            log.Info.Println("Rotating to angle ", angle)
            time.Sleep(time.Second)

            // Keep rotating
            for dist < 500 {
                direction++
                hexabody.MoveHead(direction, 0)
                dist, _ = distance.Value()
            }

            // Look up and down
            hexabody.Pitch(angle, 3 )
            log.Info.Println("Pitch to angle ", angle)
            time.Sleep(time.Second)
            // hexabody.Pitch(0 , 3 )

            moveCount++
            if (moveCount > 15) {
                moveCount = 0
                hexabody.Relax()
                time.Sleep(time.Second * 60)
                hexabody.Stand()
            }
        }

        hexabody.WalkContinuously(0, 0.5)
        time.Sleep(500 * time.Millisecond)
    }
}


func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &SensorWalkSkill{}
}

func (d *SensorWalkSkill) OnStart() {
	// Use this method to do something when this skill is starting.
    hexabody.Start()
    distance.Start()
    hexabody.Stand()
}

func (d *SensorWalkSkill) OnClose() {
	// Use this method to do something when this skill is closing.
    hexabody.Close()
    distance.Close()
}

func (d *SensorWalkSkill) OnConnect() {
	// Use this method to do something when the remote connected.
    MoveAndScan()

}

func (d *SensorWalkSkill) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
    hexabody.Relax()
}

func (d *SensorWalkSkill) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *SensorWalkSkill) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}



