package binding

import (
	"github.com/go-glx/input/system/device"
	"github.com/go-glx/input/system/input"
)

type Binding struct {
	device *device.Device
	path   input.InputPath
}

func NewBinding(device *device.Device, path input.InputPath) *Binding {
	return &Binding{
		device: device,
		path:   path,
	}
}
