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
	"a": {keybd_event.VK_A, kbEventNoFlag},
	"b": {keybd_event.VK_B, kbEventNoFlag},
	"c": {keybd_event.VK_C, kbEventNoFlag},
	"d": {keybd_event.VK_D, kbEventNoFlag},
	"e": {keybd_event.VK_E, kbEventNoFlag},
	"f": {keybd_event.VK_F, kbEventNoFlag},
	"g": {keybd_event.VK_G, kbEventNoFlag},
	"h": {keybd_event.VK_H, kbEventNoFlag},
	"i": {keybd_event.VK_I, kbEventNoFlag},
	"j": {keybd_event.VK_J, kbEventNoFlag},
	"k": {keybd_event.VK_K, kbEventNoFlag},
	"l": {keybd_event.VK_L, kbEventNoFlag},
	"m": {keybd_event.VK_M, kbEventNoFlag},
	"n": {keybd_event.VK_N, kbEventNoFlag},
	"o": {keybd_event.VK_O, kbEventNoFlag},
	"p": {keybd_event.VK_P, kbEventNoFlag},
	"q": {keybd_event.VK_Q, kbEventNoFlag},
	"r": {keybd_event.VK_R, kbEventNoFlag},
	"s": {keybd_event.VK_S, kbEventNoFlag},
	"t": {keybd_event.VK_T, kbEventNoFlag},
	"u": {keybd_event.VK_U, kbEventNoFlag},
	"v": {keybd_event.VK_V, kbEventNoFlag},
	"w": {keybd_event.VK_W, kbEventNoFlag},
	"x": {keybd_event.VK_X, kbEventNoFlag},
	"y": {keybd_event.VK_Y, kbEventNoFlag},
	"z": {keybd_event.VK_Z, kbEventNoFlag},

	"A": {keybd_event.VK_A, kbEventHasShift},
	"B": {keybd_event.VK_B, kbEventHasShift},
	"C": {keybd_event.VK_C, kbEventHasShift},
	"D": {keybd_event.VK_D, kbEventHasShift},
	"E": {keybd_event.VK_E, kbEventHasShift},
	"F": {keybd_event.VK_F, kbEventHasShift},
	"G": {keybd_event.VK_G, kbEventHasShift},
	"H": {keybd_event.VK_H, kbEventHasShift},
	"I": {keybd_event.VK_I, kbEventHasShift},
	"J": {keybd_event.VK_J, kbEventHasShift},
	"K": {keybd_event.VK_K, kbEventHasShift},
	"L": {keybd_event.VK_L, kbEventHasShift},
	"M": {keybd_event.VK_M, kbEventHasShift},
	"N": {keybd_event.VK_N, kbEventHasShift},
	"O": {keybd_event.VK_O, kbEventHasShift},
	"P": {keybd_event.VK_P, kbEventHasShift},
	"Q": {keybd_event.VK_Q, kbEventHasShift},
	"R": {keybd_event.VK_R, kbEventHasShift},
	"S": {keybd_event.VK_S, kbEventHasShift},
	"T": {keybd_event.VK_T, kbEventHasShift},
	"U": {keybd_event.VK_U, kbEventHasShift},
	"V": {keybd_event.VK_V, kbEventHasShift},
	"W": {keybd_event.VK_W, kbEventHasShift},
	"X": {keybd_event.VK_X, kbEventHasShift},
	"Y": {keybd_event.VK_Y, kbEventHasShift},
	"Z": {keybd_event.VK_Z, kbEventHasShift},

	"0": {keybd_event.VK_0, kbEventNoFlag},
	"1": {keybd_event.VK_1, kbEventNoFlag},
	"2": {keybd_event.VK_2, kbEventNoFlag},
	"3": {keybd_event.VK_3, kbEventNoFlag},
	"4": {keybd_event.VK_4, kbEventNoFlag},
	"5": {keybd_event.VK_5, kbEventNoFlag},
	"6": {keybd_event.VK_6, kbEventNoFlag},
	"7": {keybd_event.VK_7, kbEventNoFlag},
	"8": {keybd_event.VK_8, kbEventNoFlag},
	"9": {keybd_event.VK_9, kbEventNoFlag},

	")": {keybd_event.VK_0, kbEventHasShift},
	"!": {keybd_event.VK_1, kbEventHasShift},
	"@": {keybd_event.VK_2, kbEventHasShift},
	"#": {keybd_event.VK_3, kbEventHasShift},
	"$": {keybd_event.VK_4, kbEventHasShift},
	"%": {keybd_event.VK_5, kbEventHasShift},
	"^": {keybd_event.VK_6, kbEventHasShift},
	"&": {keybd_event.VK_7, kbEventHasShift},
	"*": {keybd_event.VK_8, kbEventHasShift},
	"(": {keybd_event.VK_9, kbEventHasShift},

	"-":  {keybd_event.VK_OEM_MINUS, kbEventNoFlag},
	"=":  {keybd_event.VK_OEM_PLUS, kbEventNoFlag},
	"[":  {keybd_event.VK_OEM_4, kbEventNoFlag},
	"]":  {keybd_event.VK_OEM_6, kbEventNoFlag},
	"\\": {keybd_event.VK_OEM_5, kbEventNoFlag},
	";":  {keybd_event.VK_OEM_1, kbEventNoFlag},
	"'":  {keybd_event.VK_OEM_7, kbEventNoFlag},
	",":  {keybd_event.VK_OEM_COMMA, kbEventNoFlag},
	".":  {keybd_event.VK_OEM_PERIOD, kbEventNoFlag},
	"/":  {keybd_event.VK_OEM_2, kbEventNoFlag},
	"`":  {keybd_event.VK_OEM_3, kbEventNoFlag},

	"_":  {keybd_event.VK_OEM_MINUS, kbEventHasShift},
	"+":  {keybd_event.VK_OEM_PLUS, kbEventHasShift},
	"{":  {keybd_event.VK_OEM_4, kbEventHasShift},
	"}":  {keybd_event.VK_OEM_6, kbEventHasShift},
	"|":  {keybd_event.VK_OEM_5, kbEventHasShift},
	":":  {keybd_event.VK_OEM_1, kbEventHasShift},
	"\"": {keybd_event.VK_OEM_7, kbEventHasShift},
	"<":  {keybd_event.VK_OEM_COMMA, kbEventHasShift},
	">":  {keybd_event.VK_OEM_PERIOD, kbEventHasShift},
	"?":  {keybd_event.VK_OEM_2, kbEventHasShift},
	"~":  {keybd_event.VK_OEM_3, kbEventHasShift},

	"ENTER":     {keybd_event.VK_ENTER, kbEventNoFlag},
	"TAB":       {keybd_event.VK_TAB, kbEventNoFlag},
	"BACKSPACE": {keybd_event.VK_BACK, kbEventNoFlag},
	"SPACE":     {keybd_event.VK_SPACE, kbEventNoFlag},

	"LEFT":  {keybd_event.VK_LEFT, kbEventNoFlag},
	"RIGHT": {keybd_event.VK_RIGHT, kbEventNoFlag},
	"UP":    {keybd_event.VK_UP, kbEventNoFlag},
	"DOWN":  {keybd_event.VK_DOWN, kbEventNoFlag},

	"HOME":     {keybd_event.VK_HOME, kbEventNoFlag},
	"END":      {keybd_event.VK_END, kbEventNoFlag},
	"PAGEUP":   {keybd_event.VK_PAGEUP, kbEventNoFlag},
	"PAGEDOWN": {keybd_event.VK_PAGEDOWN, kbEventNoFlag},
}
