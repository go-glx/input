package keyboard

import "github.com/go-glx/input/system/device"

type Driver interface {
	Tick()
	IsEnabled() bool
	OnEnabled(func())
	OnDisabled(func())
	IsJustPressed(key Key) bool
	IsJustReleased(key Key) bool
	IsDown(key Key) bool
}

type PhysicalDevice struct {
	driver Driver
}

func NewKeyboardDevice(driver Driver) *PhysicalDevice {
	return &PhysicalDevice{driver: driver}
}

func (v *PhysicalDevice) Kind() device.Kind {
	return device.KindKeyboard
}

func (v *PhysicalDevice) Name() device.Name {
	return "Generic virtual keyboard"
}

func (v *PhysicalDevice) IsEnabled() bool {
	return v.driver.IsEnabled()
}

func (v *PhysicalDevice) OnEnabled(fn func()) {
	v.driver.OnEnabled(fn)
}

func (v *PhysicalDevice) OnDisabled(fn func()) {
	v.driver.OnDisabled(fn)
}

func (v *PhysicalDevice) Tick() {
	v.driver.Tick()
}

func (v *PhysicalDevice) Driver() Driver {
	return v.driver
}
