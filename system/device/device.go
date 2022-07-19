package device

type Kind string
type Name string

const (
	KindMouse    Kind = "mouse"
	KindKeyboard Kind = "keyboard"
	KindGamepad  Kind = "gamepad"
	KindOther    Kind = "other"
)

type PhysicalDevice interface {
	Name() Name
	Kind() Kind
	Tick()
	IsEnabled() bool
	OnEnabled(func())
	OnDisabled(func())
}

type Device struct {
	king     Kind
	name     Name
	physical PhysicalDevice
}

func NewDevice(name Name, physical PhysicalDevice) *Device {
	return &Device{
		king:     physical.Kind(),
		name:     name,
		physical: physical,
	}
}

func (d *Device) Kind() Kind {
	return d.king
}

func (d *Device) Name() Name {
	return d.name
}

func (d *Device) Physical() PhysicalDevice {
	return d.physical
}
