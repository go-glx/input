package input

type KeyboardCustomDriver interface {
	IsEnabled() bool
	OnEnabled(func())
	OnDisabled(func())
	IsJustPressed(key KeyboardKey) bool
	IsJustReleased(key KeyboardKey) bool
	IsDown(key KeyboardKey) bool
}

type VirtualKeyboard struct {
	drv KeyboardCustomDriver
}

func NewVirtualKeyboard(drv KeyboardCustomDriver) *VirtualKeyboard {
	return &VirtualKeyboard{drv: drv}
}

func (v *VirtualKeyboard) Name() Name {
	return "Generic virtual keyboard"
}

func (v *VirtualKeyboard) Kind() Kind {
	return KindKeyboard
}

func (v *VirtualKeyboard) IsEnabled() bool {
	return v.drv.IsEnabled()
}

func (v *VirtualKeyboard) OnEnabled(fn func()) {
	v.drv.OnEnabled(fn)
}

func (v *VirtualKeyboard) OnDisabled(fn func()) {
	v.drv.OnDisabled(fn)
}

func (v *VirtualKeyboard) Driver() KeyboardCustomDriver {
	return v.drv
}
