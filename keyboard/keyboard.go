package keyboard

import "math"

const (
	alphabetMode = iota
	pinyinMode
)

type KeyboardSimulator struct {
	inputMode int
}

func (ks *KeyboardSimulator) Handle(xg *XINPUT_GAMEPAD) {
	if xg.LeftTrigger > XINPUT_GAMEPAD_TRIGGER_THRESHOLD {

	}
}

func (ks *KeyboardSimulator) judgeLeftSection(x, y int16) int {
	t := math.Atanh(float64(y) / float64(x))
	degree := int(-t+90) % 360
	interval := 360.0 / len(alphabetDict)
	return degree % interval
}

func (ks *KeyboardSimulator) changeNextMode() {
	ks.inputMode++
	if ks.inputMode >= modeCount {
		ks.inputMode = 0
	}
}

func (ks *KeyboardSimulator) changePrevMode() {
	ks.inputMode--
	if ks.inputMode <= 0 {
		ks.inputMode = modeCount - 1
	}
}
