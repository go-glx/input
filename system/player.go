package system

import "github.com/go-glx/input/system/device"

const maxPlayers = 4

const (
	PlayerUnset PlayerID = 0
	Player1     PlayerID = 1
	Player2     PlayerID = 2
	Player3     PlayerID = 3
	Player4     PlayerID = 4
)

type (
	PlayerID uint8

	Player struct {
		id         PlayerID
		active     bool
		currentMap ActionMapID
		devices    []*device.Device
	}
)

func NewPlayer(id PlayerID, active bool) *Player {
	return &Player{
		id:         id,
		active:     active,
		currentMap: "",
	}
}

// AttachDevice allows player to use this device automatically
// if player has access more than one device, first available will be used
// so provide device in AttachDevice in priority of usage
func (p *Player) AttachDevice(device *device.Device) {
	p.devices = append(p.devices, device)
}
