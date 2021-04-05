package xgc

type WCHAR uint16
type WORD uint16
type DWORD uint32
type BYTE byte
type SHORT int16

const (
	XinputGamepadLeftThumbDeadzone  = 7849
	XinputGamepadRightThumbDeadzone = 8689
	XinputGamepadTriggerThreshold   = 30
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
	Type      BYTE
	SubType   BYTE
	Flags     WORD
	Gamepad   XinputGamepad
	Vibration XinputVibration
}

type XinputGamepad struct {
	Buttons      WORD
	LeftTrigger  BYTE  // 0~255
	RightTrigger BYTE  // 0~255
	ThumbLX      SHORT // -32768~32767
	ThumbLY      SHORT // -32768~32767
	ThumbRX      SHORT // -32768~32767
	ThumbRY      SHORT // -32768~32767
}

type XinputKeyStroke struct {
	VirtualKey WORD
	Unicode    WCHAR
	Flags      WORD
	UserIndex  BYTE
	HidCode    BYTE
}

type XinputState struct {
	PacketNumber DWORD
	Gamepad      XinputGamepad
}

type XinputVibration struct {
	LeftMotorSpeed  WORD // 0~65535
	RightMotorSpeed WORD // 0~65535
}
