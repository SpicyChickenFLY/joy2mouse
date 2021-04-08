package keyboard

import (
	"math"

	"github.com/SpicyChickenFLY/xinput2mouse/xgc"
)

const (
	alphabetMode = iota
	pinyinMode

	modeCount = iota - 1
)

type KeyboardSimulator struct {
	inputMode int
}

func (ks *KeyboardSimulator) Handle(xg *xgc.XinputGamepad) {
	if xg.LeftTrigger > xgc.XinputGamepadTriggerThreshold {

	}
}

// judgePosSection 计算所处第几个扇形（以圆形正上方为起点）
// x, y: 横纵坐标（代表摇杆对应的百分比）
// l: 代表该圆中存在几个扇形
func (ks *KeyboardSimulator) judgePosSection(x, y float64, l int) int {
	temp := math.Atanh(y / x)
	degree := int(-temp+90) % 360 // 将极坐标正方向反转后逆时针旋转90度（求余用于溢出角度）
	interval := 360.0 / l
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
