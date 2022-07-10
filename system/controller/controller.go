package controller

import (
	"github.com/google/uuid"

	"github.com/go-glx/input/system/device"
)

type ID string

// Controller is list of virtual devices
// that can be used for resolving input
// One player can use only one controller at same time
//
// Controller is useful for fast switching control schema
// for one player, or for making hot-seat multiplier
type Controller struct {
	id      ID
	devices []*device.Device
}

func NewController(devices []*device.Device) *Controller {
	return &Controller{
		id:      ID(uuid.NewString()),
		devices: devices,
	}
}

func (ctl *Controller) ID() ID {
	return ctl.id
}

func (ctl *Controller) Devices() []*device.Device {
	return ctl.devices
}
