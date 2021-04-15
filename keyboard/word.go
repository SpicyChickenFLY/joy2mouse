package keyboard

import "github.com/micmonay/keybd_event"

var alphabetLoweraseDict = [][]string{
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "s", "r"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
	{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
	{"-", "=", "[", "]", "\\", ";", "'", ",", ".", "/", "`"},
}

var alphabetHigherCaseDict = [][]string{
	{"A", "B", "C"},
	{"D", "E", "F"},
	{"G", "H", "I"},
	{"J", "K", "L"},
	{"M", "N", "O"},
	{"P", "Q", "S", "R"},
	{"T", "U", "V"},
	{"W", "X", "Y", "Z"},
	{")", "!", "@", "#", "$", "%", "^", "&", "*", "("},
	{"_", "+", "{", "}", "|", ":", "\"", "<", ">", "?", "~"},
}

var pinyinDict = [][][]string{
	{
		{"b", "p", "m", "f"},
		{"d", "t", "l", "n"},
		{"g", "k", "g"},
		{"j", "q", "x"},
		{"z", "c", "s"},
		{"zh", "ch", "sh"},
		{"y", "w"},
	},
	{{"a", "o", "e", "i", "u", "v"}},
}

const (
	kbEventNoFlag   = 0x0000
	kbEventHasShift = 0x0001
	kbEventHasCtrl  = 0x0002
	kbEventHasAlt   = 0x0004
	kvEventHasWin   = 0x0008
)

type keyInfo struct {
	VK   int
	Flag int
}

var keyInfoMap = map[string]keyInfo{
	"a": keyInfo{keybd_event.VK_A, kbEventNoFlag},
	"b": keyInfo{keybd_event.VK_B, kbEventNoFlag},
	"c": keyInfo{keybd_event.VK_C, kbEventNoFlag},
	"d": keyInfo{keybd_event.VK_D, kbEventNoFlag},
	"e": keyInfo{keybd_event.VK_E, kbEventNoFlag},
	"f": keyInfo{keybd_event.VK_F, kbEventNoFlag},
	"g": keyInfo{keybd_event.VK_G, kbEventNoFlag},
	"h": keyInfo{keybd_event.VK_H, kbEventNoFlag},
	"i": keyInfo{keybd_event.VK_I, kbEventNoFlag},
	"j": keyInfo{keybd_event.VK_J, kbEventNoFlag},
	"k": keyInfo{keybd_event.VK_K, kbEventNoFlag},
	"l": keyInfo{keybd_event.VK_L, kbEventNoFlag},
	"m": keyInfo{keybd_event.VK_M, kbEventNoFlag},
	"n": keyInfo{keybd_event.VK_N, kbEventNoFlag},
	"o": keyInfo{keybd_event.VK_O, kbEventNoFlag},
	"p": keyInfo{keybd_event.VK_P, kbEventNoFlag},
	"q": keyInfo{keybd_event.VK_Q, kbEventNoFlag},
	"r": keyInfo{keybd_event.VK_R, kbEventNoFlag},
	"s": keyInfo{keybd_event.VK_S, kbEventNoFlag},
	"t": keyInfo{keybd_event.VK_T, kbEventNoFlag},
	"u": keyInfo{keybd_event.VK_U, kbEventNoFlag},
	"v": keyInfo{keybd_event.VK_V, kbEventNoFlag},
	"w": keyInfo{keybd_event.VK_W, kbEventNoFlag},
	"x": keyInfo{keybd_event.VK_X, kbEventNoFlag},
	"y": keyInfo{keybd_event.VK_Y, kbEventNoFlag},
	"z": keyInfo{keybd_event.VK_Z, kbEventNoFlag},

	"A": keyInfo{keybd_event.VK_A, kbEventHasShift},
	"B": keyInfo{keybd_event.VK_B, kbEventHasShift},
	"C": keyInfo{keybd_event.VK_C, kbEventHasShift},
	"D": keyInfo{keybd_event.VK_D, kbEventHasShift},
	"E": keyInfo{keybd_event.VK_E, kbEventHasShift},
	"F": keyInfo{keybd_event.VK_F, kbEventHasShift},
	"G": keyInfo{keybd_event.VK_G, kbEventHasShift},
	"H": keyInfo{keybd_event.VK_H, kbEventHasShift},
	"I": keyInfo{keybd_event.VK_I, kbEventHasShift},
	"J": keyInfo{keybd_event.VK_J, kbEventHasShift},
	"K": keyInfo{keybd_event.VK_K, kbEventHasShift},
	"L": keyInfo{keybd_event.VK_L, kbEventHasShift},
	"M": keyInfo{keybd_event.VK_M, kbEventHasShift},
	"N": keyInfo{keybd_event.VK_N, kbEventHasShift},
	"O": keyInfo{keybd_event.VK_O, kbEventHasShift},
	"P": keyInfo{keybd_event.VK_P, kbEventHasShift},
	"Q": keyInfo{keybd_event.VK_Q, kbEventHasShift},
	"R": keyInfo{keybd_event.VK_R, kbEventHasShift},
	"S": keyInfo{keybd_event.VK_S, kbEventHasShift},
	"T": keyInfo{keybd_event.VK_T, kbEventHasShift},
	"U": keyInfo{keybd_event.VK_U, kbEventHasShift},
	"V": keyInfo{keybd_event.VK_V, kbEventHasShift},
	"W": keyInfo{keybd_event.VK_W, kbEventHasShift},
	"X": keyInfo{keybd_event.VK_X, kbEventHasShift},
	"Y": keyInfo{keybd_event.VK_Y, kbEventHasShift},
	"Z": keyInfo{keybd_event.VK_Z, kbEventHasShift},

	"0": keyInfo{keybd_event.VK_0, kbEventNoFlag},
	"1": keyInfo{keybd_event.VK_1, kbEventNoFlag},
	"2": keyInfo{keybd_event.VK_2, kbEventNoFlag},
	"3": keyInfo{keybd_event.VK_3, kbEventNoFlag},
	"4": keyInfo{keybd_event.VK_4, kbEventNoFlag},
	"5": keyInfo{keybd_event.VK_5, kbEventNoFlag},
	"6": keyInfo{keybd_event.VK_6, kbEventNoFlag},
	"7": keyInfo{keybd_event.VK_7, kbEventNoFlag},
	"8": keyInfo{keybd_event.VK_8, kbEventNoFlag},
	"9": keyInfo{keybd_event.VK_9, kbEventNoFlag},

	")": keyInfo{keybd_event.VK_0, kbEventHasShift},
	"!": keyInfo{keybd_event.VK_1, kbEventHasShift},
	"@": keyInfo{keybd_event.VK_2, kbEventHasShift},
	"#": keyInfo{keybd_event.VK_3, kbEventHasShift},
	"$": keyInfo{keybd_event.VK_4, kbEventHasShift},
	"%": keyInfo{keybd_event.VK_5, kbEventHasShift},
	"^": keyInfo{keybd_event.VK_6, kbEventHasShift},
	"&": keyInfo{keybd_event.VK_7, kbEventHasShift},
	"*": keyInfo{keybd_event.VK_8, kbEventHasShift},
	"(": keyInfo{keybd_event.VK_9, kbEventHasShift},

	"-":  keyInfo{keybd_event.VK_OEM_MINUS, kbEventNoFlag},
	"=":  keyInfo{keybd_event.VK_OEM_PLUS, kbEventNoFlag},
	"[":  keyInfo{keybd_event.VK_OEM_4, kbEventNoFlag},
	"]":  keyInfo{keybd_event.VK_OEM_6, kbEventNoFlag},
	"\\": keyInfo{keybd_event.VK_OEM_5, kbEventNoFlag},
	";":  keyInfo{keybd_event.VK_OEM_1, kbEventNoFlag},
	"'":  keyInfo{keybd_event.VK_OEM_7, kbEventNoFlag},
	",":  keyInfo{keybd_event.VK_OEM_COMMA, kbEventNoFlag},
	".":  keyInfo{keybd_event.VK_OEM_PERIOD, kbEventNoFlag},
	"/":  keyInfo{keybd_event.VK_OEM_2, kbEventNoFlag},
	"`":  keyInfo{keybd_event.VK_OEM_3, kbEventNoFlag},

	"_":  keyInfo{keybd_event.VK_OEM_MINUS, kbEventHasShift},
	"+":  keyInfo{keybd_event.VK_OEM_PLUS, kbEventHasShift},
	"{":  keyInfo{keybd_event.VK_OEM_4, kbEventHasShift},
	"}":  keyInfo{keybd_event.VK_OEM_6, kbEventHasShift},
	"|":  keyInfo{keybd_event.VK_OEM_5, kbEventHasShift},
	":":  keyInfo{keybd_event.VK_OEM_1, kbEventHasShift},
	"\"": keyInfo{keybd_event.VK_OEM_7, kbEventHasShift},
	"<":  keyInfo{keybd_event.VK_OEM_COMMA, kbEventHasShift},
	">":  keyInfo{keybd_event.VK_OEM_PERIOD, kbEventHasShift},
	"?":  keyInfo{keybd_event.VK_OEM_2, kbEventHasShift},
	"~":  keyInfo{keybd_event.VK_OEM_3, kbEventHasShift},

	"ENTER":     keyInfo{keybd_event.VK_ENTER, kbEventNoFlag},
	"TAB":       keyInfo{keybd_event.VK_TAB, kbEventNoFlag},
	"BACKSPACE": keyInfo{keybd_event.VK_BACK, kbEventNoFlag},
	"SPACE":     keyInfo{keybd_event.VK_SPACE, kbEventNoFlag},
}
