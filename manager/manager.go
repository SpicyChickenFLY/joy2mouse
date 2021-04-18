package manager

import (
	"errors"

	"github.com/SpicyChickenFLY/xinput2mouse/keyboard"
	"github.com/SpicyChickenFLY/xinput2mouse/mouse"
	"github.com/SpicyChickenFLY/xinput2mouse/screen"
	"github.com/SpicyChickenFLY/xinput2mouse/xgc"
)

// screen mode cant be called by LB/RB
const (
	joystickMode = -1 - iota
	screenMode
)

const (
	mouseMode = iota
	keyboardMode

	modeCount
)

// Manager handle xbox game controller input and deliver
type Manager struct {
	mode          int
	joystick      xgc.Joystick
	lastPacketNum uint32
	kbSim         *keyboard.Simulator
	mSim          *mouse.Simulator
	screen        *screen.Screen
}

// NewManager init new manager
func NewManager() Manager {
	return Manager{
		joystick: xgc.NewXGC(0),
		kbSim:    keyboard.NewSimulator(),
	}
}

// Loop
func (m *Manager) Loop() error {
	for true {
		if err := m.handleEvent(); err != nil {
			return err
		}
	}
	return nil
}

// handleEvent handle event
func (m *Manager) handleEvent() error {
	stateInter, err := m.joystick.GetState()
	if err != nil {
		return err
	}
	state := stateInter.(xgc.XinputState)
	if state.PacketNumber == m.lastPacketNum {
		return nil
	}

	// FIXME: 需要知道xinput的拦截方式，和重新发送的方法
	if m.mode == joystickMode {
		// only interrupted when RB-LB-ThumbL-ThumbR pushed both
		if state.Gamepad.Buttons&
			(xgc.XinputGamepadLeftShoulder|
				xgc.XinputGamepadRightShoulder|
				xgc.XinputGamepadLeftThumb|
				xgc.XinputGamepadRightThumb) > 0 {
			m.mode = mouseMode
		}
	} else {
		switch { // if changing mode, ignore any other input
		case (state.Gamepad.Buttons & xgc.XinputGamepadLeftShoulder) > 0:
			m.changePrevMode()
			return nil
		case (state.Gamepad.Buttons & xgc.XinputGamepadRightShoulder) > 0:
			m.changeNextMode()
			return nil
		case (state.Gamepad.Buttons & xgc.XinputGamepadStart) > 0:
			// m.changeScreenMode()
			// return nil
			panic(errors.New("DEBUG"))
		}

		switch m.mode { // Deliver event to correspond simulator
		case mouseMode:
			m.mSim.Handle(&state.Gamepad)
		case keyboardMode:
			m.kbSim.Handle(&state.Gamepad)
		case screenMode:
			m.screen.Handle(&state.Gamepad)
		}
	}

	return nil
}

func (m *Manager) changeScreenMode() {
	if m.mode == screenMode {
		m.mode = mouseMode
	} else {
		m.mode = screenMode
	}
}

func (m *Manager) changeNextMode() {
	m.mode++
	if m.mode >= modeCount {
		m.mode = 0
	}
}

func (m *Manager) changePrevMode() {
	m.mode--
	if m.mode <= 0 {
		m.mode = modeCount - 1
	}
}
