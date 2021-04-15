package keyboard

import (
	"math"

	"github.com/SpicyChickenFLY/xinput2mouse/xgc"
	"github.com/micmonay/keybd_event"
)

const (
	alphabetMode = iota
	// pinyinMode

	modeCount = iota - 1
)

// judgePosSection 计算所处第几个扇形（以圆形正上方为起点）
// x, y: 横纵坐标（代表摇杆对应的百分比）
// l: 代表该圆中存在几个扇形
func judgePosSection(x, y float64, l int) int {
	temp := math.Atanh(y / x)
	degree := int(-temp+90) % 360 // 将极坐标正方向反转后逆时针旋转90度（求余用于溢出角度）
	interval := 360.0 / l
	return degree % interval
}

type KeyboardSimulator struct {
	inputMode int
	kbEvent   keybd_event.KeyBonding
}

func (ks *KeyboardSimulator) InitKB(xg *xgc.XinputGamepad) (err error) {
	ks.kbEvent, err = keybd_event.NewKeyBonding()
	return err
}

func (ks *KeyboardSimulator) Handle(xg *xgc.XinputGamepad) error {
	var ltPulled bool
	var alphabetDict *[][]string
	// First judge if Left Trigger pulled
	if xg.JudgeLTPulled() {
		ltPulled = true
		alphabetDict = &alphabetHigherCaseDict
	} else {
		ltPulled = false
		alphabetDict = &alphabetLoweraseDict
	}
	// if Right Trigger pulled, judge which key should be simulated
	if xg.JudgeRTPulled() {
		lSec, rSec := 0, 0
		if xg.JudgeThumbLPulled() {
			lx := float64(xg.ThumbLX) / xgc.ThumbMax
			ly := float64(xg.ThumbLY) / xgc.ThumbMax
			ll := len(*alphabetDict)
			lSec = judgePosSection(lx, ly, ll)
		}
		if lSec != 0 {
			if xg.JudgeThumbRPulled() {
				rx := float64(xg.ThumbRX) / xgc.ThumbMax
				ry := float64(xg.ThumbRY) / xgc.ThumbMax
				rl := len((*alphabetDict)[lSec])
				rSec = judgePosSection(rx, ry, rl)
			}
		}
		if rSec != 0 {
			keyVal := (*alphabetDict)[lSec][rSec]
			ks.kbEvent.HasSHIFT(keyInfoMap[keyVal].Flag&kbEventHasShift > 0)
			ks.kbEvent.AddKey(keyInfoMap[keyVal].VK)
			err := ks.kbEvent.Launching()
			if err != nil {
				return err
			}
		}
	}
	// judge rest input in order and not aborted
	if xg.Buttons&xgc.XinputGamepadDpad > 0 {
		if ltPulled {

		} else {

		}
	}
	return nil
}

// func (ks *KeyboardSimulator)
