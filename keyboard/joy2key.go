package keyboard

import "github.com/SpicyChickenFLY/xinput2mouse/xgc"

var dpadKeyDict = [2]map[uint16]string{
	{
		xgc.XinputGamepadDpadUp:    "PAGEUP",
		xgc.XinputGamepadDpadDown:  "PAGEDOWN",
		xgc.XinputGamepadDpadLeft:  "HOME",
		xgc.XinputGamepadDpadRight: "END",
	},
	{
		xgc.XinputGamepadDpadUp:    "UP",
		xgc.XinputGamepadDpadDown:  "DOWN",
		xgc.XinputGamepadDpadLeft:  "LEFT",
		xgc.XinputGamepadDpadRight: "RIGHT",
	},
}

var mainKeyDict = [2]map[uint16]string{
	{
		xgc.XinputGamepadY: "TAB",
	},
	{
		xgc.XinputGamepadA: "SPACE",
		xgc.XinputGamepadB: "ENTER",
		xgc.XinputGamepadX: "BACKSPACE",
		xgc.XinputGamepadY: "TAB",
	},
}
