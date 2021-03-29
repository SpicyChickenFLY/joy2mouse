package manager

import (
	"github.com/SpicyChickenFLY/joy2mouse/xgc"
)

const (
	mouseMode = iota
	keyboardMode
)
const modeCount = 2

type Manager struct {
	joyMode  int
	joystick xgc.Joystick
}

func NewManager() Manager {
	return Manager{}
}

// func (m *Manager) HandleEvent(event, value int) {
// 	if event == XINPUT_GAMEPAD_LEFT_SHOULDER {
// 		m.joyMode
// 	}
// }

func (m *Manager) changeNextMode() {
	m.joyMode++
	if m.joyMode >= modeCount {
		m.joyMode = 0
	}
}

func (m *Manager) changePrevMode() {
	m.joyMode--
	if m.joyMode <= 0 {
		m.joyMode = modeCount - 1
	}
}
