package device

type Kind string
type Name string

const (
	KindMouse    Kind = "mouse"
	KindKeyboard Kind = "keyboard"
	KindGamepad  Kind = "gamepad"
	KindOther    Kind = "other"
)

type Device struct {
	king  Kind  // kind of device
	index uint8 // connected device index (for example if we have 2 gamepads)
	name  Name  // some human readable device name (like "xbox controller")
}

func NewDevice(kind Kind, index uint8, name Name) *Device {
	return &Device{
		king:  kind,
		index: index,
		name:  name,
	}
}
