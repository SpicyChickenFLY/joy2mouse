package event

import "github.com/micmonay/keybd_event"

const (
	KBEventNoFlag   = 0x0000
	KBEventHasShift = 0x0001
	KBEventHasCtrl  = 0x0002
	KBEventHasAlt   = 0x0004
	KBEventHasWin   = 0x0008
)

type keyInfo struct {
	VK   int
	Flag int
}

var KeyInfoMap = map[string]keyInfo{
	"a": {keybd_event.VK_A, KBEventNoFlag},
	"b": {keybd_event.VK_B, KBEventNoFlag},
	"c": {keybd_event.VK_C, KBEventNoFlag},
	"d": {keybd_event.VK_D, KBEventNoFlag},
	"e": {keybd_event.VK_E, KBEventNoFlag},
	"f": {keybd_event.VK_F, KBEventNoFlag},
	"g": {keybd_event.VK_G, KBEventNoFlag},
	"h": {keybd_event.VK_H, KBEventNoFlag},
	"i": {keybd_event.VK_I, KBEventNoFlag},
	"j": {keybd_event.VK_J, KBEventNoFlag},
	"k": {keybd_event.VK_K, KBEventNoFlag},
	"l": {keybd_event.VK_L, KBEventNoFlag},
	"m": {keybd_event.VK_M, KBEventNoFlag},
	"n": {keybd_event.VK_N, KBEventNoFlag},
	"o": {keybd_event.VK_O, KBEventNoFlag},
	"p": {keybd_event.VK_P, KBEventNoFlag},
	"q": {keybd_event.VK_Q, KBEventNoFlag},
	"r": {keybd_event.VK_R, KBEventNoFlag},
	"s": {keybd_event.VK_S, KBEventNoFlag},
	"t": {keybd_event.VK_T, KBEventNoFlag},
	"u": {keybd_event.VK_U, KBEventNoFlag},
	"v": {keybd_event.VK_V, KBEventNoFlag},
	"w": {keybd_event.VK_W, KBEventNoFlag},
	"x": {keybd_event.VK_X, KBEventNoFlag},
	"y": {keybd_event.VK_Y, KBEventNoFlag},
	"z": {keybd_event.VK_Z, KBEventNoFlag},

	"A": {keybd_event.VK_A, KBEventHasShift},
	"B": {keybd_event.VK_B, KBEventHasShift},
	"C": {keybd_event.VK_C, KBEventHasShift},
	"D": {keybd_event.VK_D, KBEventHasShift},
	"E": {keybd_event.VK_E, KBEventHasShift},
	"F": {keybd_event.VK_F, KBEventHasShift},
	"G": {keybd_event.VK_G, KBEventHasShift},
	"H": {keybd_event.VK_H, KBEventHasShift},
	"I": {keybd_event.VK_I, KBEventHasShift},
	"J": {keybd_event.VK_J, KBEventHasShift},
	"K": {keybd_event.VK_K, KBEventHasShift},
	"L": {keybd_event.VK_L, KBEventHasShift},
	"M": {keybd_event.VK_M, KBEventHasShift},
	"N": {keybd_event.VK_N, KBEventHasShift},
	"O": {keybd_event.VK_O, KBEventHasShift},
	"P": {keybd_event.VK_P, KBEventHasShift},
	"Q": {keybd_event.VK_Q, KBEventHasShift},
	"R": {keybd_event.VK_R, KBEventHasShift},
	"S": {keybd_event.VK_S, KBEventHasShift},
	"T": {keybd_event.VK_T, KBEventHasShift},
	"U": {keybd_event.VK_U, KBEventHasShift},
	"V": {keybd_event.VK_V, KBEventHasShift},
	"W": {keybd_event.VK_W, KBEventHasShift},
	"X": {keybd_event.VK_X, KBEventHasShift},
	"Y": {keybd_event.VK_Y, KBEventHasShift},
	"Z": {keybd_event.VK_Z, KBEventHasShift},

	"0": {keybd_event.VK_0, KBEventNoFlag},
	"1": {keybd_event.VK_1, KBEventNoFlag},
	"2": {keybd_event.VK_2, KBEventNoFlag},
	"3": {keybd_event.VK_3, KBEventNoFlag},
	"4": {keybd_event.VK_4, KBEventNoFlag},
	"5": {keybd_event.VK_5, KBEventNoFlag},
	"6": {keybd_event.VK_6, KBEventNoFlag},
	"7": {keybd_event.VK_7, KBEventNoFlag},
	"8": {keybd_event.VK_8, KBEventNoFlag},
	"9": {keybd_event.VK_9, KBEventNoFlag},

	")": {keybd_event.VK_0, KBEventHasShift},
	"!": {keybd_event.VK_1, KBEventHasShift},
	"@": {keybd_event.VK_2, KBEventHasShift},
	"#": {keybd_event.VK_3, KBEventHasShift},
	"$": {keybd_event.VK_4, KBEventHasShift},
	"%": {keybd_event.VK_5, KBEventHasShift},
	"^": {keybd_event.VK_6, KBEventHasShift},
	"&": {keybd_event.VK_7, KBEventHasShift},
	"*": {keybd_event.VK_8, KBEventHasShift},
	"(": {keybd_event.VK_9, KBEventHasShift},

	"-":  {keybd_event.VK_OEM_MINUS, KBEventNoFlag},
	"=":  {keybd_event.VK_OEM_PLUS, KBEventNoFlag},
	"[":  {keybd_event.VK_OEM_4, KBEventNoFlag},
	"]":  {keybd_event.VK_OEM_6, KBEventNoFlag},
	"\\": {keybd_event.VK_OEM_5, KBEventNoFlag},
	";":  {keybd_event.VK_OEM_1, KBEventNoFlag},
	"'":  {keybd_event.VK_OEM_7, KBEventNoFlag},
	",":  {keybd_event.VK_OEM_COMMA, KBEventNoFlag},
	".":  {keybd_event.VK_OEM_PERIOD, KBEventNoFlag},
	"/":  {keybd_event.VK_OEM_2, KBEventNoFlag},
	"`":  {keybd_event.VK_OEM_3, KBEventNoFlag},

	"_":  {keybd_event.VK_OEM_MINUS, KBEventHasShift},
	"+":  {keybd_event.VK_OEM_PLUS, KBEventHasShift},
	"{":  {keybd_event.VK_OEM_4, KBEventHasShift},
	"}":  {keybd_event.VK_OEM_6, KBEventHasShift},
	"|":  {keybd_event.VK_OEM_5, KBEventHasShift},
	":":  {keybd_event.VK_OEM_1, KBEventHasShift},
	"\"": {keybd_event.VK_OEM_7, KBEventHasShift},
	"<":  {keybd_event.VK_OEM_COMMA, KBEventHasShift},
	">":  {keybd_event.VK_OEM_PERIOD, KBEventHasShift},
	"?":  {keybd_event.VK_OEM_2, KBEventHasShift},
	"~":  {keybd_event.VK_OEM_3, KBEventHasShift},

	"ENTER":     {keybd_event.VK_ENTER, KBEventNoFlag},
	"TAB":       {keybd_event.VK_TAB, KBEventNoFlag},
	"BACKSPACE": {keybd_event.VK_BACK, KBEventNoFlag},
	"SPACE":     {keybd_event.VK_SPACE, KBEventNoFlag},

	"LEFT":  {keybd_event.VK_LEFT, KBEventNoFlag},
	"RIGHT": {keybd_event.VK_RIGHT, KBEventNoFlag},
	"UP":    {keybd_event.VK_UP, KBEventNoFlag},
	"DOWN":  {keybd_event.VK_DOWN, KBEventNoFlag},

	"HOME":     {keybd_event.VK_HOME, KBEventNoFlag},
	"END":      {keybd_event.VK_END, KBEventNoFlag},
	"PAGEUP":   {keybd_event.VK_PAGEUP, KBEventNoFlag},
	"PAGEDOWN": {keybd_event.VK_PAGEDOWN, KBEventNoFlag},
}
