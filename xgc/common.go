package xgc

// type WCHAR uint16
// type WORD uint16
// type DWORD uint32
// type BYTE byte
// type SHORT int16

const (
	XinputGamepadLeftThumbDeadzone  = 7849
	XinputGamepadRightThumbDeadzone = 8689
	XinputGamepadTriggerThreshold   = 30
)

const (
	ThumbMax = 32767
)

const (
	XinputGamepadDpadUp        = 0x0001
	XinputGamepadDpadDown      = 0x0002
	XinputGamepadDpadLeft      = 0x0004
	XinputGamepadDpadRight     = 0x0008
	XinputGamepadStart         = 0x0010
	XinputGamepadBack          = 0x0020
	XinputGamepadLeftThumb     = 0x0040
	XinputGamepadRightThumb    = 0x0080
	XinputGamepadLeftShoulder  = 0x0100
	XinputGamepadRightShoulder = 0x0200
	XinputGamepadA             = 0x1000
	XinputGamepadB             = 0x2000
	XinputGamepadX             = 0x4000
	XinputGamepadY             = 0x8000
)

// -+-CAPABILITIES
//  |-----GAMEPAD
//  |-----VIBRATION
//  +-KEYSTROKE
//  +-STATE
//  |-----GAMEPAD

type XinputCapabilities struct {
	Type      byte
	SubType   byte
	Flags     uint16
	Gamepad   XinputGamepad
	Vibration XinputVibration
}

type XinputGamepad struct {
	Buttons      uint16
	LeftTrigger  byte  // 0~255
	RightTrigger byte  // 0~255
	ThumbLX      int16 // -32768~32767
	ThumbLY      int16 // -32768~32767
	ThumbRX      int16 // -32768~32767
	ThumbRY      int16 // -32768~32767
}

func (xg *XinputGamepad) JudgeLTPulled() bool {
	return xg.LeftTrigger > XinputGamepadTriggerThreshold
}

func (xg *XinputGamepad) JudgeRTPulled() bool {
	return xg.RightTrigger > XinputGamepadTriggerThreshold
}

func (xg *XinputGamepad) JudgeThumbLPulled() bool {
	if xg.ThumbLX > XinputGamepadLeftThumbDeadzone {
		return true
	}
	if xg.ThumbLY > XinputGamepadLeftThumbDeadzone {
		return true
	}
	return false
}

func (xg *XinputGamepad) JudgeThumbRPulled() bool {
	if xg.ThumbRX > XinputGamepadRightThumbDeadzone {
		return true
	}
	if xg.ThumbRY > XinputGamepadRightThumbDeadzone {
		return true
	}
	return false
}

type XinputKeyStroke struct {
	VirtualKey uint16
	Unicode    uint16
	Flags      uint16
	UserIndex  byte
	HidCode    byte
}

type XinputState struct {
	PacketNumber uint32
	Gamepad      XinputGamepad
}

type XinputVibration struct {
	LeftMotorSpeed  uint16 // 0~65535
	RightMotorSpeed uint16 // 0~65535
}
