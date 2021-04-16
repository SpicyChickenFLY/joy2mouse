package manager

import (
	"github.com/SpicyChickenFLY/xinput2mouse/keyboard"
	"github.com/SpicyChickenFLY/xinput2mouse/mouse"
	"github.com/SpicyChickenFLY/xinput2mouse/xgc"
)

// screen mode cant be called by LB/RB
const screenMode = -1

const (
	joystickMode = iota
	mouseMode
	keyboardMode

	modeCount
)

type Manager struct {
	joyMode       int
	joystick      xgc.Joystick
	lastPacketNum uint32
	kbSim         *keyboard.Simulator
	mSim          *mouse.Simulator
}

func NewManager() Manager {
	return Manager{
		kbSim: keyboard.NewSimulator(),
	}
}

func (m *Manager) HandleEvent(event, value int) error {
	stateInter, err := m.joystick.GetState()
	if err != nil {
		return err
	}
	state := stateInter.(xgc.XinputState)
	if state.PacketNumber == m.lastPacketNum {
		return nil
	}
	// if changing mode, ignore any other input
	if (state.Gamepad.Buttons & xgc.XinputGamepadLeftShoulder) > 0 {
		m.changePrevMode()
		return nil
	}
	if (state.Gamepad.Buttons & xgc.XinputGamepadRightShoulder) > 0 {
		m.changeNextMode()
		return nil
	}

	// if pushed any func buttons(start,back), ignore any other input
	if (state.Gamepad.Buttons & xgc.XinputGamepadBack) > 0 {

		return nil
	}
	if (state.Gamepad.Buttons & xgc.XinputGamepadStart) > 0 {

		return nil
	}

	// Deliver to correspond simulator
	switch m.joyMode {
	case mouseMode:

	case keyboardMode:
		m.kbSim.Handle(&state.Gamepad)
	case joystickMode:

	}

	return nil
}

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
