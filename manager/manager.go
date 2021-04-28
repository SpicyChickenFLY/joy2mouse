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
	mode      int
	joystick  xgc.Joystick
	currState *xgc.XinputState
	lastState *xgc.XinputState
	// button
	kbSim  *keyboard.Simulator
	mSim   *mouse.Simulator
	screen *screen.Screen
}

// NewManager init new manager
func NewManager() Manager {
	return Manager{
		joystick: xgc.NewXGC(0),
		kbSim:    keyboard.NewSimulator(),
	}
}

// Loop to update Joystick until
func (m *Manager) Loop() error {
	for true {
		// FIXME: 最后应当要保证
		if err := m.updateJoystick(); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) updateJoystick() error {
	// 检查手柄状态有无变化
	stateInter, err := m.joystick.GetState()
	if err != nil {
		return err
	}
	state := stateInter.(xgc.XinputState)
	if state.PacketNumber == m.lastState.PacketNumber {
		return nil
	}
	// FIXME: 是否需要做防震动处理（按下和松开曲线不平稳）
	switch m.mode {
	case joystickMode:
		// only interrupted when RB-LB-ThumbL-ThumbR pushed both
		if m.judgeBtnsHold(xgc.XinputGamepadLeftShoulder |
			xgc.XinputGamepadRightShoulder |
			xgc.XinputGamepadLeftThumb |
			xgc.XinputGamepadRightThumb) {
			m.mode = mouseMode
		}
	case mouseMode:
		m.mSim.Handle(&state.Gamepad)
	case keyboardMode:
		m.kbSim.Update(&state.Gamepad)
	case screenMode:
		m.screen.Handle(&state.Gamepad)
	}
	// Traverse all button flags
	for eventIdx := xgc.XinputGamepadDpadUp; eventIdx <= xgc.XinputGamepadY; eventIdx <<= 1 {
		if m.judgeBtnUp(uint16(eventIdx)) {
			m.handleEvent(uint16(eventIdx))
		}
	}
	m.lastState = &state
	return nil
}

func (m *Manager) handleEvent(eventIdx uint16) error {

	switch eventIdx { // if changing mode, ignore any other input
	case xgc.XinputGamepadLeftShoulder:
		m.prevMode()
		return nil
	case xgc.XinputGamepadRightShoulder:
		m.nextMode()
		return nil
	case xgc.XinputGamepadStart:
		// m.changeScreenMode()
		// return nil
		panic(errors.New("DEBUG"))
	}

	switch m.mode { // Deliver event to correspond simulator
	case mouseMode:
		// return m.mSim.Handle(&state.Gamepad)
	case keyboardMode:
		return m.kbSim.HandleEvent(eventIdx)
	case screenMode:
		// return m.screen.Handle(&state.Gamepad)
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

func (m *Manager) nextMode() {
	m.mode++
	if m.mode >= modeCount {
		m.mode = 0
	}
}

func (m *Manager) prevMode() {
	m.mode--
	if m.mode <= 0 {
		m.mode = modeCount - 1
	}
}

func (m *Manager) judgeBtnsHold(buttonFlags uint16) bool {
	return buttonFlags&m.currState.Gamepad.Buttons == buttonFlags
}

func (m *Manager) judgeBtnUp(buttonFlags uint16) bool {
	return buttonFlags&m.lastState.Gamepad.Buttons > 0 && buttonFlags&m.currState.Gamepad.Buttons == 0
}
